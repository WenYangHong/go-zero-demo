<template>
  <main class="max-w-7xl mx-auto px-6 py-8">
    <template v-if="loading">
      <div class="grid grid-cols-4 gap-3 mb-8">
        <div class="col-span-4 md:col-span-2 row-span-2">
          <el-skeleton animated>
            <template #template>
              <el-skeleton-item variant="image" style="width: 100%; height: 100%; min-height: 280px; border-radius: 0.75rem;" />
            </template>
          </el-skeleton>
        </div>
        <div v-for="i in 4" :key="i">
          <el-skeleton animated>
            <template #template>
              <el-skeleton-item variant="image" style="width: 100%; height: 140px; border-radius: 0.75rem;" />
            </template>
          </el-skeleton>
        </div>
      </div>

      <div class="grid grid-cols-1 md:grid-cols-3 gap-8 mb-8">
        <div class="md:col-span-2">
          <el-skeleton animated :rows="2">
            <template #template>
              <div class="flex items-center gap-2 mb-3">
                <el-skeleton-item variant="text" style="width: 80px; height: 28px; border-radius: 9999px;" />
                <el-skeleton-item variant="text" style="width: 72px; height: 28px; border-radius: 9999px;" />
              </div>
              <el-skeleton-item variant="h1" style="width: 50%; height: 32px;" />
              <el-skeleton-item variant="text" style="width: 70%; margin-top: 8px;" />

              <div class="flex items-center gap-6 py-4 border-y border-outline-variant/30 my-6">
                <el-skeleton-item variant="text" style="width: 120px;" />
                <el-skeleton-item variant="text" style="width: 120px;" />
              </div>

              <div class="bg-surface rounded-xl p-5 shadow-card mb-6">
                <div class="flex items-center gap-4">
                  <el-skeleton-item variant="circle" style="width: 56px; height: 56px;" />
                  <div class="flex-1">
                    <el-skeleton-item variant="text" style="width: 40%;" />
                    <el-skeleton-item variant="text" style="width: 60%; margin-top: 8px;" />
                  </div>
                </div>
              </div>

              <div class="mb-6">
                <el-skeleton-item variant="h3" style="width: 100px; height: 24px; margin-bottom: 16px;" />
                <el-skeleton-item variant="text" style="width: 100%;" />
                <el-skeleton-item variant="text" style="width: 90%; margin-top: 8px;" />
                <el-skeleton-item variant="text" style="width: 95%; margin-top: 8px;" />
              </div>

              <div class="mb-6">
                <el-skeleton-item variant="h3" style="width: 100px; height: 24px; margin-bottom: 16px;" />
                <div class="grid grid-cols-3 md:grid-cols-4 gap-4">
                  <el-skeleton-item v-for="i in 8" :key="i" variant="text" style="width: 80px;" />
                </div>
              </div>

              <div class="mb-6">
                <el-skeleton-item variant="h3" style="width: 100px; height: 24px; margin-bottom: 16px;" />
                <div class="bg-surface rounded-xl p-5 shadow-card">
                  <div class="flex justify-between mb-3">
                    <el-skeleton-item variant="text" style="width: 80px;" />
                    <el-skeleton-item variant="text" style="width: 60px;" />
                  </div>
                  <el-skeleton-item variant="text" style="width: 100%;" />
                </div>
              </div>

              <div>
                <div class="flex justify-between mb-4">
                  <el-skeleton-item variant="h3" style="width: 80px; height: 24px;" />
                  <el-skeleton-item variant="text" style="width: 60px;" />
                </div>
                <div class="space-y-4">
                  <div v-for="i in 2" :key="i" class="bg-surface rounded-xl p-5 shadow-card">
                    <div class="flex items-start gap-4">
                      <el-skeleton-item variant="circle" style="width: 40px; height: 40px;" />
                      <div class="flex-1">
                        <el-skeleton-item variant="text" style="width: 60px;" />
                        <el-skeleton-item variant="text" style="width: 100%; margin-top: 12px;" />
                        <el-skeleton-item variant="text" style="width: 80%; margin-top: 8px;" />
                      </div>
                    </div>
                  </div>
                </div>
              </div>
            </template>
          </el-skeleton>
        </div>

        <div>
          <div class="bg-surface rounded-xl shadow-card p-6">
            <el-skeleton animated :rows="5">
              <template #template>
                <el-skeleton-item variant="h1" style="width: 40%; height: 36px; margin-bottom: 16px;" />
                <div class="grid grid-cols-2 gap-3 mb-4">
                  <el-skeleton-item variant="rect" style="width: 100%; height: 40px; border-radius: 0.5rem;" />
                  <el-skeleton-item variant="rect" style="width: 100%; height: 40px; border-radius: 0.5rem;" />
                </div>
                <el-skeleton-item variant="rect" style="width: 100%; height: 40px; border-radius: 0.5rem; margin-bottom: 16px;" />
                <el-skeleton-item variant="rect" style="width: 100%; height: 40px; border-radius: 0.5rem; margin-bottom: 8px;" />
                <el-skeleton-item variant="rect" style="width: 100%; height: 40px; border-radius: 0.5rem;" />
              </template>
            </el-skeleton>
          </div>
        </div>
      </div>
    </template>

    <template v-else-if="isEmpty">
      <div class="flex flex-col items-center justify-center py-24">
        <div class="w-32 h-32 bg-surface-container rounded-3xl flex items-center justify-center mb-8">
          <el-icon :size="64" class="text-on-surface-variant/30"><HomeFilled /></el-icon>
        </div>
        <h2 class="text-xl font-bold text-on-surface mb-2">民宿已下架或不存在</h2>
        <p class="text-on-surface-variant mb-8">该民宿可能已被房东下架，或链接地址有误</p>
        <div class="flex gap-3">
          <el-button @click="router.back()">
            <el-icon class="mr-2"><ArrowLeft /></el-icon>
            返回上一页
          </el-button>
          <el-button type="primary" @click="router.push('/')">
            <el-icon class="mr-2"><HomeFilled /></el-icon>
            回到首页
          </el-button>
        </div>
      </div>
    </template>

    <template v-else>
      <div class="grid grid-cols-4 gap-3 mb-8">
        <template v-if="homestay.images.length > 0">
          <div class="col-span-4 md:col-span-2 row-span-2">
            <img
              :src="homestay.images[0]"
              alt="民宿主图"
              class="w-full h-full object-cover rounded-xl"
              @error="handleImgError"
            />
          </div>
          <div v-for="(img, idx) in homestay.images.slice(1, 5)" :key="idx">
            <img
              :src="img"
              :alt="`民宿图片${idx + 2}`"
              class="w-full h-full object-cover rounded-xl"
              @error="handleImgError"
            />
          </div>
        </template>
        <template v-else>
          <div class="col-span-4 bg-surface-container rounded-xl flex items-center justify-center" style="height: 280px;">
            <span class="text-on-surface-variant">暂无图片</span>
          </div>
        </template>
      </div>

      <div class="grid grid-cols-1 md:grid-cols-3 gap-8 mb-8">
        <div class="md:col-span-2">
          <div class="flex items-start justify-between mb-4">
            <div>
              <div class="flex items-center gap-2 mb-2">
                <span class="px-3 py-1 bg-primary/10 text-primary text-sm font-medium rounded-full">{{ homestay.category }}</span>
                <span class="px-3 py-1 bg-surface-container text-on-surface-variant text-sm rounded-full">可住{{ homestay.capacity }}人</span>
              </div>
              <h1 class="text-2xl font-bold text-on-surface mb-2">{{ homestay.title }}</h1>
              <p v-if="homestay.subtitle" class="text-on-surface-variant">{{ homestay.subtitle }}</p>
            </div>
            <span class="flex items-center gap-1 text-sm text-on-surface-variant">
              <span class="text-primary font-bold line-through">¥{{ homestay.marketPrice }}</span>
            </span>
          </div>

          <div class="flex items-center gap-6 py-4 border-y border-outline-variant/30 mb-6">
            <div class="flex items-center gap-2">
              <el-icon class="text-warning" :size="20"><StarFilled /></el-icon>
              <span class="font-semibold">{{ homestay.rating || '暂无' }}</span>
              <span class="text-on-surface-variant">({{ homestay.reviewCount }}条评价)</span>
            </div>
            <div v-if="homestay.location" class="flex items-center gap-2">
              <el-icon class="text-primary" :size="20"><Location /></el-icon>
              <span class="text-on-surface-variant">{{ homestay.location }}</span>
            </div>
          </div>

          <div v-if="boss.name" class="bg-surface rounded-xl p-5 shadow-card mb-6">
            <div class="flex items-center gap-4">
              <el-avatar :size="56" :src="boss.avatar" class="bg-primary/10 text-primary text-lg font-bold">
                {{ boss.name.charAt(0) }}
              </el-avatar>
              <div class="flex-1">
                <h3 class="font-semibold text-on-surface">房东：{{ boss.name }}</h3>
                <p class="text-sm text-on-surface-variant">{{ boss.info }}</p>
              </div>
              <router-link :to="`/business/${boss.id}`" class="text-primary text-sm font-medium hover:underline">
                查看主页
              </router-link>
            </div>
          </div>

          <div class="mb-6">
            <h2 class="text-lg font-bold text-on-surface mb-4">民宿介绍</h2>
            <p v-for="(p, idx) in homestay.description" :key="idx" class="text-on-surface-variant leading-relaxed" :class="{ 'mt-3': idx > 0 }">
              {{ p }}
            </p>
          </div>

          <div class="mb-6">
            <h2 class="text-lg font-bold text-on-surface mb-4">设施服务</h2>
            <div class="grid grid-cols-3 md:grid-cols-4 gap-4">
              <div v-for="facility in homestay.facilities" :key="facility.name" class="flex items-center gap-3 text-sm">
                <el-icon class="text-primary" :size="20"><component :is="facility.icon" /></el-icon>
                <span>{{ facility.name }}</span>
              </div>
            </div>
          </div>

          <div class="mb-6">
            <h2 class="text-lg font-bold text-on-surface mb-4">餐食服务</h2>
            <div class="bg-surface rounded-xl p-5 shadow-card">
              <div class="flex items-center justify-between mb-3">
                <span class="font-medium">{{ homestay.meal.name }}</span>
                <span class="text-primary font-bold">¥{{ homestay.meal.price }}/人</span>
              </div>
              <p class="text-sm text-on-surface-variant">{{ homestay.meal.description }}</p>
            </div>
          </div>

          <div>
            <div class="flex items-center justify-between mb-4">
              <h2 class="text-lg font-bold text-on-surface">用户评价</h2>
              <router-link :to="`/comment/${homestay.id}`" class="text-primary text-sm font-medium hover:underline">
                查看全部
              </router-link>
            </div>
            <div v-if="comments.length > 0" class="space-y-4">
              <div v-for="review in comments" :key="review.id" class="bg-surface rounded-xl p-5 shadow-card">
                <div class="flex items-start gap-4">
                  <el-avatar :size="40" class="bg-primary/10 text-primary font-bold shrink-0">
                    {{ review.author.charAt(0) }}
                  </el-avatar>
                  <div class="flex-1">
                    <div class="flex items-center gap-2 mb-2">
                      <span class="font-medium">{{ review.author }}</span>
                      <el-rate v-model="review.rating" disabled :colors="['#f59e0b', '#f59e0b', '#f59e0b']" size="small" />
                    </div>
                    <p class="text-sm text-on-surface-variant leading-relaxed">{{ review.content }}</p>
                    <p class="text-xs text-on-surface-variant mt-2">{{ review.date }}</p>
                  </div>
                </div>
              </div>
            </div>
            <div v-else class="bg-surface rounded-xl p-10 shadow-card text-center text-on-surface-variant">
              暂无评价，成为第一个评价的人吧
            </div>
          </div>
        </div>

        <div>
          <div class="bg-surface rounded-xl shadow-card p-6 sticky top-24">
            <div class="flex items-baseline gap-2 mb-4">
              <span class="text-3xl font-bold text-primary">¥{{ homestay.price }}</span>
              <span class="text-on-surface-variant">/晚</span>
            </div>

            <div class="space-y-4 mb-6 booking-form">
              <div class="grid grid-cols-2 gap-3">
                <div>
                  <label class="block text-xs text-on-surface-variant mb-1">入住</label>
                  <el-date-picker
                    v-model="bookingForm.checkIn"
                    type="date"
                    placeholder="选择入住"
                    size="large"
                    class="w-full"
                    :disabled-date="disablePastDate"
                    value-format="YYYY-MM-DD"
                  />
                </div>
                <div>
                  <label class="block text-xs text-on-surface-variant mb-1">退房</label>
                  <el-date-picker
                    v-model="bookingForm.checkOut"
                    type="date"
                    placeholder="选择退房"
                    size="large"
                    class="w-full"
                    :disabled-date="disableCheckOutDate"
                    value-format="YYYY-MM-DD"
                  />
                </div>
              </div>
              <div>
                <label class="block text-xs text-on-surface-variant mb-1">人数</label>
                <el-select v-model="bookingForm.guestCount" size="large" class="w-full">
                  <el-option
                    v-for="n in homestay.capacity"
                    :key="n"
                    :label="`${n}人`"
                    :value="n"
                  />
                </el-select>
              </div>
            </div>

            <div class="flex flex-col gap-3">
              <el-button type="primary" size="large" class="w-full" @click="handleBook">
                立即预订
              </el-button>
              <el-button size="large" class="w-full contact-house-master-btn" @click="handleContact">
                联系房东
              </el-button>
            </div>

            <div class="mt-6 pt-6 border-t border-outline-variant/30">
              <div class="flex justify-between text-sm mb-2">
                <span class="text-on-surface-variant">¥{{ homestay.price }} x {{ nights }}晚</span>
                <span>¥{{ homestay.price * nights }}</span>
              </div>
              <div class="flex justify-between text-sm mb-2">
                <span class="text-on-surface-variant">清洁费</span>
                <span>¥50</span>
              </div>
              <div class="flex justify-between text-sm mb-2">
                <span class="text-on-surface-variant">服务费</span>
                <span>¥68</span>
              </div>
              <div class="flex justify-between font-semibold pt-3 border-t border-outline-variant/30 mt-3">
                <span>总计</span>
                <span class="text-primary">¥{{ totalPrice }}</span>
              </div>
            </div>
          </div>
        </div>
      </div>
    </template>
  </main>
</template>

<script setup>
import { ref, reactive, computed, markRaw, onMounted } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { ElMessage } from 'element-plus'
import {
  StarFilled, Location, Connection, Van, IceCreamRound,
  Sunny, Coffee, Grid, House, KnifeFork,
  HomeFilled, ArrowLeft
} from '@element-plus/icons-vue'
import { getHomestayDetail, getHomestayBossWithHomestayId, getHomestayCommentList } from '@/api/homestay.js'

const route = useRoute()
const router = useRouter()

const loading = ref(true);
const isEmpty = ref(false);

const bookingForm = reactive({
  checkIn: '',
  checkOut: '',
  guestCount: 2,
})

const DEFAULT_IMAGES = [
  'https://images.unsplash.com/photo-1600596542815-ffad4c1539a9?w=800&h=600&fit=crop',
  'https://images.unsplash.com/photo-1600585154340-be6161a56a0c?w=400&h=300&fit=crop',
  'https://images.unsplash.com/photo-1600607687939-ce8a6c25118c?w=400&h=300&fit=crop',
  'https://images.unsplash.com/photo-1600566753190-17f0baa2a6c3?w=400&h=300&fit=crop',
  'https://images.unsplash.com/photo-1600573472591-ee6b68d14c68?w=400&h=300&fit=crop',
]

const DEFAULT_FACILITIES = [
  { name: '高速WiFi', icon: markRaw(Connection) },
  { name: '免费停车', icon: markRaw(Van) },
  { name: '空调', icon: markRaw(IceCreamRound) },
  { name: '地暖', icon: markRaw(Sunny) },
  { name: '24小时热水', icon: markRaw(Coffee) },
  { name: '独立厨房', icon: markRaw(Grid) },
  { name: '花园庭院', icon: markRaw(House) },
  { name: '烧烤设备', icon: markRaw(KnifeFork) },
]

function trim(s) {
  return (s || '').replace(/^\n+|\n+$/g, '').trim()
}

const homestay = ref({
  id: '',
  title: '',
  subtitle: '',
  category: '精品民宿',
  capacity: 0,
  rating: 0,
  reviewCount: 0,
  location: '',
  price: 0,
  marketPrice: 0,
  images: [],
  description: [],
  facilities: [...DEFAULT_FACILITIES],
  meal: {
    name: '暂无餐食',
    price: 0,
    description: '',
  },
})

const boss = ref({
  id: '',
  name: '',
  info: '',
  cover: '',
  avatar: '',
})

const comments = ref([])

const nights = computed(() => {
  if (!bookingForm.checkIn || !bookingForm.checkOut) return 1
  const diff = new Date(bookingForm.checkOut) - new Date(bookingForm.checkIn)
  const days = Math.ceil(diff / 86400000)
  return days > 0 ? days : 1
})

const totalPrice = computed(() => {
  const price = homestay.value.price || 0
  return price * nights.value + 50 + 68
})

function disablePastDate(date) {
  return date.getTime() < Date.now() - 86400000
}

function disableCheckOutDate(date) {
  if (!bookingForm.checkIn) return date.getTime() < Date.now() - 86400000
  return date.getTime() <= new Date(bookingForm.checkIn).getTime()
}

function handleBook() {
  if (!bookingForm.checkIn || !bookingForm.checkOut) {
    ElMessage.warning('请选择入住和退房日期')
    return
  }
  ElMessage.success('预订功能开发中，敬请期待')
}

function handleContact() {
  ElMessage.info('联系房东功能开发中，敬请期待')
}

function handleImgError(e) {
  e.target.src = 'data:image/svg+xml,' + encodeURIComponent(
    `<svg xmlns="http://www.w3.org/2000/svg" width="400" height="300"><rect width="100%" height="100%" fill="#f3f4f6"/><text x="50%" y="50%" text-anchor="middle" font-size="24" fill="#9ca3af" dominant-baseline="central">🖼</text></svg>`
  )
}

async function fetchHomestayDetail() {
  const res = await getHomestayDetail({ id: route.params.id })

  if (res.code === 200002) {
    isEmpty.value = true
    return
  }

  const h = res.homestay
  const images = h.img_info && h.img_info.length > 0 ? h.img_info : DEFAULT_IMAGES

  homestay.value = {
    id: h.id,
    title: trim(h.title),
    subtitle: trim(h.sub_title),
    category: trim(h.tag) || '精品民宿',
    capacity: h.people_num || 0,
    rating: h.comment_avg_score || 0,
    reviewCount: h.comment_count || 0,
    location: '',
    price: h.homestay_price || 0,
    marketPrice: h.market_homestay_price || 0,
    images,
    description: trim(h.info)
      .split('\n\n')
      .map(p => p.trim())
      .filter(Boolean),
    facilities: [...DEFAULT_FACILITIES],
    meal: {
      name: trim(h.food_info) || '暂无餐食',
      price: h.food_price || 0,
      description: trim(h.food_info) || '',
    },
  }

  if (homestay.value.capacity > 0) {
    bookingForm.guestCount = Math.min(2, homestay.value.capacity)
  }
}

async function fetchBoss() {
  try {
    const res = await getHomestayBossWithHomestayId({ homestay_id: route.params.id })
    boss.value = {
      id: res.id,
      name: trim(res.title),
      info: trim(res.boss_info),
      cover: res.cover,
      avatar: res.HeaderImg,
    }
  } catch {
    boss.value = { id: '', name: '', info: '', cover: '', avatar: '' }
  }
}

async function fetchComments() {
  try {
    const res = await getHomestayCommentList({
      id: route.params.id,
      page: 1,
      page_size: 2,
    })
    comments.value = (res.list || []).map((item, index) => ({
      id: index,
      author: item.user_name,
      rating: item.avg_star_result,
      content: item.content,
      date: item.create_time,
    }))
  } catch {
    comments.value = []
  }
}

onMounted(async () => {
  try {
    await fetchHomestayDetail()
    if (!isEmpty.value) {
      await Promise.all([
        fetchBoss(),
        fetchComments(),
      ])
    }
  } finally {
    loading.value = false
  }
})
</script>

<style scoped>
.booking-form :deep(.el-date-editor) {
  width: 100% !important;
}

.booking-form :deep(.el-input__wrapper) {
  background-color: var(--color-surface-container);
  box-shadow: none;
  border-radius: 0.5rem;
  padding: 0 0.75rem;
}

.booking-form :deep(.el-input__wrapper:hover) {
  box-shadow: none;
}

.booking-form :deep(.el-input__wrapper.is-focus) {
  box-shadow: 0 0 0 2px rgba(5, 150, 105, 0.3);
}

.booking-form :deep(.el-select .el-input__wrapper) {
  background-color: var(--color-surface-container);
  border-radius: 0.5rem;
}

.contact-house-master-btn {
  margin-left: 0;
}
</style>
