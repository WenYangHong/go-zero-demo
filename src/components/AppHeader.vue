<template>
  <header class="bg-surface sticky top-0 z-40 border-b border-outline-variant/30">
    <div class="max-w-7xl mx-auto px-6 h-16 flex items-center justify-between">
      <router-link to="/" class="flex items-center gap-2">
        <el-icon :size="24" class="text-primary"><HomeFilled /></el-icon>
        <span class="font-bold text-xl text-on-surface">归栖民宿</span>
      </router-link>
      <nav class="flex items-center gap-1">
        <router-link
          to="/"
          class="px-4 py-2 text-sm font-medium rounded-lg transition-colors"
          :class="isActive('/') ? 'text-primary bg-primary/10' : 'text-on-surface-variant hover:text-on-surface hover:bg-surface-container'"
        >
          发现民宿
        </router-link>
        <router-link
          to="/?tab=hot"
          class="px-4 py-2 text-sm font-medium rounded-lg transition-colors"
          :class="isActive('/?tab=hot') ? 'text-primary bg-primary/10' : 'text-on-surface-variant hover:text-on-surface hover:bg-surface-container'"
        >
          热门推荐
        </router-link>
        <router-link
          to="/?tab=host"
          class="px-4 py-2 text-sm font-medium rounded-lg transition-colors"
          :class="isActive('/?tab=host') ? 'text-primary bg-primary/10' : 'text-on-surface-variant hover:text-on-surface hover:bg-surface-container'"
        >
          特色房东
        </router-link>
      </nav>
      <div class="flex items-center gap-3">
        <template v-if="userStore.isLoggedIn">
          <el-dropdown trigger="click">
            <span class="flex items-center gap-2 cursor-pointer">
              <el-avatar :size="32" class="bg-primary/10 text-primary">
                {{ userStore.nickname?.charAt(0) || '用' }}
              </el-avatar>
            </span>
            <template #dropdown>
              <el-dropdown-menu>
                <el-dropdown-item @click="handleLogout">退出登录</el-dropdown-item>
              </el-dropdown-menu>
            </template>
          </el-dropdown>
        </template>
        <template v-else>
          <router-link to="/login" class="px-4 py-2 text-sm font-medium text-on-surface hover:text-primary transition-colors">
            登录
          </router-link>
          <router-link to="/register" class="px-4 py-2 text-sm font-medium bg-primary text-on-primary rounded-lg hover:opacity-90 transition-opacity">
            注册
          </router-link>
        </template>
      </div>
    </div>
  </header>
</template>

<script setup>
import { useRoute, useRouter } from 'vue-router'
import { HomeFilled } from '@element-plus/icons-vue'
import { useUserStore } from '@/stores/user'

const route = useRoute()
const router = useRouter()
const userStore = useUserStore()

function isActive(path) {
  if (path === '/') return route.path === '/' && !route.query.tab
  return route.fullPath === path
}

async function handleLogout() {
  await userStore.logout()
  router.push('/')
}
</script>
