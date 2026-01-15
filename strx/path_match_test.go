package strx

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func ExamplePathMatch() {
	fmt.Println(PathMatch("/api/users", "/api/users"))
	fmt.Println(PathMatch("/api/users*", "/api/users123"))
	fmt.Println(PathMatch("/api/users*", "/api/users/add"))
	fmt.Println(PathMatch("/api/**/detail", "/api/user/info/detail"))
	fmt.Println(PathMatch("/api/**", "/api"))

	// Output:
	// true
	// true
	// false
	// true
	// false
}

func Test_MatchPath(t *testing.T) {
	tests := []struct {
		name    string
		pattern string
		path    string
		want    bool
	}{
		// ========= 精确匹配 =========
		{
			name:    "精确路径完全匹配",
			pattern: "/api/users",
			path:    "/api/users",
			want:    true,
		},
		{
			name:    "精确路径不匹配",
			pattern: "/api/users",
			path:    "/api/user",
			want:    false,
		},

		// ========= 段内 * 通配 =========
		{
			name:    "段内前缀通配 users* 可以匹配同层路径",
			pattern: "/api/users*",
			path:    "/api/users666",
			want:    true,
		},
		{
			name:    "段内前缀通配 users* 不能跨层级匹配",
			pattern: "/api/users*",
			path:    "/api/users/add",
			want:    false,
		},
		{
			name:    "段内后缀通配 *Test 可以匹配同层路径",
			pattern: "/api/*Test",
			path:    "/api/UserTest",
			want:    true,
		},
		{
			name:    "段内前后缀组合通配 user*Test 可以匹配",
			pattern: "/api/user*Test",
			path:    "/api/userABC123Test",
			want:    true,
		},

		// ========= 单层 * =========
		{
			name:    "单层通配符 * 匹配一个路径段",
			pattern: "/api/*/detail",
			path:    "/api/user/detail",
			want:    true,
		},
		{
			name:    "单层通配符 * 不能匹配多个路径段",
			pattern: "/api/*/detail",
			path:    "/api/user/info/detail",
			want:    false,
		},

		// ========= 多层 ** =========
		{
			name:    "多层通配符 ** 匹配多个路径段",
			pattern: "/api/**",
			path:    "/api/user/info/detail",
			want:    true,
		},
		{
			name:    "多层通配符 ** 必须至少匹配一层路径",
			pattern: "/api/**",
			path:    "/api",
			want:    false,
		},
		{
			name:    "多层通配符 ** 位于中间可以匹配多层路径",
			pattern: "/api/**/detail",
			path:    "/api/user/info/detail",
			want:    true,
		},
		{
			name:    "多层通配符 ** 位于中间但不能匹配零层",
			pattern: "/api/**/detail",
			path:    "/api/detail",
			want:    false,
		},

		// ========= 组合规则 =========
		{
			name:    "单层通配与段内通配组合匹配成功",
			pattern: "/api/*/users*",
			path:    "/api/admin/usersAdd",
			want:    true,
		},
		{
			name:    "单层通配与段内通配组合匹配失败",
			pattern: "/api/*/users*",
			path:    "/api/admin/users/add",
			want:    false,
		},
		{
			name:    "多层通配与段内通配组合匹配成功",
			pattern: "/api/**/users*",
			path:    "/api/a/b/c/usersTest",
			want:    true,
		},

		// ========= 路径规范化 =========
		{
			name:    "权限路径缺少前导斜杠仍可匹配",
			pattern: "api/users",
			path:    "/api/users",
			want:    true,
		},
		{
			name:    "权限路径包含尾部斜杠仍可匹配",
			pattern: "/api/users/",
			path:    "/api/users",
			want:    true,
		},

		// ========= 边界与反例 =========
		{
			name:    "权限路径层级多于请求路径时不匹配",
			pattern: "/api/users/add",
			path:    "/api/users",
			want:    false,
		},
		{
			name:    "请求路径层级多于权限路径且无通配符时不匹配",
			pattern: "/api/users",
			path:    "/api/users/add",
			want:    false,
		},
		{
			name:    "根路径不能匹配非根路径",
			pattern: "/",
			path:    "/api",
			want:    false,
		},
		{
			name:    "根路径可以精确匹配根路径",
			pattern: "/",
			path:    "/",
			want:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ok := PathMatch(tt.pattern, tt.path)
			require.Equal(t, tt.want, ok)
		})
	}
}

func Benchmark_MatchPath(b *testing.B) {
	pattern := "/api/**/users*"
	path := "/api/admin/system/v1/usersDetail"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = PathMatch(pattern, path)
	}
}
