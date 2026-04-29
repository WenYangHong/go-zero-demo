<template>
  <footer class="bg-surface border-t border-outline-variant/30 mt-16">
    <div class="max-w-7xl mx-auto px-6 py-12">
      <div class="grid grid-cols-4 gap-8">
        <div>
          <div class="flex items-center gap-2 mb-4">
            <el-icon :size="20" class="text-primary"><HomeFilled /></el-icon>
            <span class="font-bold text-lg">归栖民宿</span>
          </div>
          <p class="text-sm text-on-surface-variant leading-relaxed">
            精选各地特色民宿，<br>为您打造温馨舒适的旅途之家。
          </p>
        </div>
        <div>
          <h4 class="font-semibold text-sm mb-4">关于我们</h4>
          <ul class="space-y-2 text-sm text-on-surface-variant">
            <li><a href="#" class="hover:text-primary transition-colors">平台介绍</a></li>
            <li><a href="#" class="hover:text-primary transition-colors">加入我们</a></li>
            <li><a href="#" class="hover:text-primary transition-colors">联系方式</a></li>
          </ul>
        </div>
        <div>
          <h4 class="font-semibold text-sm mb-4">帮助中心</h4>
          <ul class="space-y-2 text-sm text-on-surface-variant">
            <li><a href="#" class="hover:text-primary transition-colors">常见问题</a></li>
            <li><a href="#" class="hover:text-primary transition-colors">预订指南</a></li>
            <li><a href="#" class="hover:text-primary transition-colors">退款政策</a></li>
          </ul>
        </div>
        <div>
          <h4 class="font-semibold text-sm mb-4">关注我们</h4>
          <div class="flex gap-3">
            <a href="#" class="w-9 h-9 rounded-lg bg-surface-container flex items-center justify-center hover:bg-primary hover:text-on-primary transition-colors">
              <el-icon :size="16"><Promotion /></el-icon>
            </a>
            <a href="#" class="w-9 h-9 rounded-lg bg-surface-container flex items-center justify-center hover:bg-primary hover:text-on-primary transition-colors">
              <el-icon :size="16"><Iphone /></el-icon>
            </a>
            <a href="#" class="w-9 h-9 rounded-lg bg-surface-container flex items-center justify-center hover:bg-primary hover:text-on-primary transition-colors">
              <el-icon :size="16"><VideoCamera /></el-icon>
            </a>
          </div>
        </div>
      </div>
      <div class="mt-8 pt-8 border-t border-outline-variant/30 text-center text-sm text-on-surface-variant">
        © 2024 归栖民宿 All Rights Reserved
      </div>
    </div>
  </footer>
</template>

<script setup>
import { HomeFilled, Promotion, Iphone, VideoCamera } from '@element-plus/icons-vue'
</script>
