<template>
  <router-link :to="`/homestay/${homestay.id}`" class="group bg-surface rounded-xl shadow-card overflow-hidden hover:shadow-float transition-all block">
    <div class="relative aspect-[4/3] overflow-hidden">
      <img
        :src="homestay.banner"
        :alt="homestay.title"
        class="w-full h-full object-cover group-hover:scale-105 transition-transform duration-300"
        @error="handleImgError"
      />
      {{  homestay }}
      <span
        v-if="homestay.tag"
        class="absolute top-3 left-3 px-3 py-1 text-xs font-medium rounded-full text-white"
        :style="{ backgroundColor: tagColor }"
      >
        {{ homestay.tag }}
      </span>
    </div>
    <div class="p-4">
      <div class="flex items-center gap-2 mb-2">
        <span class="px-2 py-0.5 bg-primary/10 text-primary text-xs rounded">{{ homestay.tag }}</span>
        <span class="px-2 py-0.5 bg-surface-container text-on-surface-variant text-xs rounded">可住{{ homestay.people_num }}人</span>
      </div>
      <h3 class="font-semibold text-on-surface mb-1 group-hover:text-primary transition-colors">{{ homestay.title }}</h3>
      <p class="text-sm text-on-surface-variant mb-3 line-clamp-2">{{ homestay.sub_title }}</p>
      <div class="flex items-center justify-between">
        <div class="flex items-center gap-1">
          <el-icon class="text-warning"><StarFilled /></el-icon>
          <span class="text-sm font-medium">{{ homestay.comment_avg_score }}</span>
          <span class="text-sm text-on-surface-variant">({{ homestay.comment_count }}条评价)</span>
        </div>
        <div>
          <span class="text-lg font-bold text-primary">¥{{ homestay.homestay_price }}</span>
          <span class="text-sm text-on-surface-variant">/晚</span>
        </div>
      </div>
    </div>
  </router-link>
</template>

<script setup>
import { computed } from 'vue'
import { StarFilled } from '@element-plus/icons-vue'

const props = defineProps({
  homestay: {
    type: Object,
    required: true,
  },
})

const tagColorMap = {
  '精选': '#059669',
  '海景': '#3b82f6',
  '古镇': '#f59e0b',
  '森林': '#16a34a',
  '精品': '#059669',
  '露营': '#f97316',
  '温泉': '#3b82f6',
}

const tagColor = computed(() => tagColorMap[props.homestay.tag] || '#059669')

function handleImgError(e) {
  e.target.src = 'data:image/svg+xml,' + encodeURIComponent(
    `<svg xmlns="http://www.w3.org/2000/svg" width="400" height="300"><rect width="100%" height="100%" fill="#f3f4f6"/><text x="50%" y="50%" text-anchor="middle" font-size="24" fill="#9ca3af" dominant-baseline="central">🖼</text></svg>`
  )
}
</script>
