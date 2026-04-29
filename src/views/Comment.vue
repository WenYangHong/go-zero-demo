<template>
  <main class="max-w-4xl mx-auto px-6 py-8">
    <div class="bg-surface rounded-xl p-6 shadow-card mb-8">
      <div class="flex gap-4">
        <img
          :src="homestayInfo.cover"
          alt="民宿图片"
          class="w-32 h-24 object-cover rounded-lg"
          @error="handleImgError"
        />
        <div class="flex-1">
          <h2 class="font-semibold text-on-surface mb-1">{{ homestayInfo.title }}</h2>
          <p class="text-sm text-on-surface-variant mb-2">{{ homestayInfo.location }} | {{ homestayInfo.checkInDate }}入住</p>
          <div class="flex items-center gap-2">
            <el-icon class="text-warning" :size="16"><StarFilled /></el-icon>
            <span class="text-sm font-medium">{{ homestayInfo.rating }}</span>
            <span class="text-sm text-on-surface-variant">房东：{{ homestayInfo.hostName }}</span>
          </div>
        </div>
      </div>
    </div>

    <div class="bg-surface rounded-xl p-6 shadow-card mb-8">
      <h2 class="text-lg font-bold text-on-surface mb-6">我要评价</h2>

      <div class="space-y-6">
        <div class="bg-surface-container rounded-xl p-5">
          <div class="flex items-center justify-between">
            <label class="text-sm font-medium text-on-surface">总体评分</label>
            <span v-if="form.overallScore > 0" class="text-lg font-bold" :class="scoreColorClass">
              {{ form.overallScore }} 分
            </span>
            <span v-else class="text-sm text-on-surface-variant">点击星星评分</span>
          </div>
          <div class="flex items-center gap-3 mt-3">
            <el-rate
              v-model="form.overallScore"
              :size="36"
              :colors="['#f59e0b', '#f59e0b', '#f59e0b']"
            />
          </div>
        </div>

        <div class="grid grid-cols-2 md:grid-cols-4 gap-3">
          <div
            v-for="dim in dimensions"
            :key="dim.key"
            class="bg-surface-container rounded-xl p-4 flex flex-col items-center"
          >
            <label class="text-xs font-medium text-on-surface-variant mb-2">{{ dim.label }}</label>
            <el-rate
              v-model="form[dim.key]"
              size="small"
              :colors="['#f59e0b', '#f59e0b', '#f59e0b']"
            />
            <span v-if="form[dim.key] > 0" class="mt-1 text-xs font-medium text-primary">{{ form[dim.key] }}分</span>
            <span v-else class="mt-1 text-xs text-outline">未评分</span>
          </div>
        </div>

        <div>
          <label class="block text-sm font-medium text-on-surface mb-2">评价内容</label>
          <el-input
            v-model="form.content"
            type="textarea"
            :rows="5"
            placeholder="分享您的住宿体验吧~（10-500字）"
            maxlength="500"
            show-word-limit
          />
        </div>

        <div>
          <label class="block text-sm font-medium text-on-surface mb-3">上传图片（选填，最多9张）</label>
          <el-upload
            v-model:file-list="form.images"
            action="#"
            list-type="picture-card"
            :auto-upload="false"
            :limit="9"
            accept="image/*"
          >
            <el-icon :size="28"><Plus /></el-icon>
          </el-upload>
        </div>

        <div class="flex items-center gap-2">
          <el-checkbox v-model="form.anonymous" label="匿名评价" />
        </div>
      </div>

      <div class="mt-6 pt-6 border-t border-outline-variant/30 flex gap-3">
        <el-button type="primary" size="large" @click="handleSubmit">
          提交评价
        </el-button>
        <el-button size="large" @click="handleCancel">
          取消
        </el-button>
      </div>
    </div>

    <div class="bg-surface-container rounded-xl p-6">
      <h3 class="font-semibold text-on-surface mb-3">评价小贴士</h3>
      <ul class="space-y-2 text-sm text-on-surface-variant">
        <li class="flex items-start gap-2">
          <el-icon class="text-primary shrink-0 mt-0.5" :size="16"><CircleCheckFilled /></el-icon>
          <span>真实的评价能帮助其他旅客做出更好的选择</span>
        </li>
        <li class="flex items-start gap-2">
          <el-icon class="text-primary shrink-0 mt-0.5" :size="16"><CircleCheckFilled /></el-icon>
          <span>分享住宿的具体体验，如房间设施、周边环境等</span>
        </li>
        <li class="flex items-start gap-2">
          <el-icon class="text-primary shrink-0 mt-0.5" :size="16"><CircleCheckFilled /></el-icon>
          <span>上传真实拍摄的图片，让评价更有说服力</span>
        </li>
      </ul>
    </div>
  </main>
</template>

<script setup>
import { reactive, ref, computed } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { ElMessage } from 'element-plus'
import { StarFilled, Plus, CircleCheckFilled } from '@element-plus/icons-vue'

const route = useRoute()
const router = useRouter()

const rateLabels = {
  1: '非常差',
  2: '比较差',
  3: '一般',
  4: '很好',
  5: '非常好',
}

const dimensions = [
  { key: 'cleanliness', label: '清洁程度' },
  { key: 'location', label: '位置交通' },
  { key: 'service', label: '服务态度' },
  { key: 'value', label: '性价比' },
]

const form = reactive({
  overallScore: 0,
  cleanliness: 0,
  location: 0,
  service: 0,
  value: 0,
  content: '',
  images: [],
  anonymous: false,
})

const scoreColorClass = computed(() => {
  const s = form.overallScore
  if (s === 0) return 'text-on-surface-variant'
  if (s <= 2) return 'text-error'
  if (s <= 3) return 'text-warning'
  return 'text-primary'
})

const homestayInfo = ref({
  id: route.params.homestayId || 1,
  title: '云顶山居·观星木屋',
  cover: 'https://images.unsplash.com/photo-1600596542815-ffad4c1539a9?w=200&h=150&fit=crop',
  location: '浙江省·湖州市',
  checkInDate: '2024年3月15日',
  rating: 4.9,
  hostName: '林志远',
})

function handleSubmit() {
  if (form.overallScore === 0) {
    ElMessage.warning('请选择总体评分')
    return
  }
  if (!form.content || form.content.length < 10) {
    ElMessage.warning('评价内容至少10个字')
    return
  }
  ElMessage.success('评价提交成功！感谢您的反馈。')
  router.push('/')
}

function handleCancel() {
  router.push('/')
}

function handleImgError(e) {
  e.target.src = 'data:image/svg+xml,' + encodeURIComponent(
    `<svg xmlns="http://www.w3.org/2000/svg" width="400" height="300"><rect width="100%" height="100%" fill="#f3f4f6"/><text x="50%" y="50%" text-anchor="middle" font-size="24" fill="#9ca3af" dominant-baseline="central">🖼</text></svg>`
  )
}
</script>
