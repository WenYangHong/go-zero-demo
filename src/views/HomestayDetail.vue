<template>
  <main class="max-w-7xl mx-auto px-6 py-8">
    <div class="grid grid-cols-4 gap-3 mb-8">
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
            <p class="text-on-surface-variant">{{ homestay.subtitle }}</p>
          </div>
        </div>

        <div class="flex items-center gap-6 py-4 border-y border-outline-variant/30 mb-6">
          <div class="flex items-center gap-2">
            <el-icon class="text-warning" :size="20"><StarFilled /></el-icon>
            <span class="font-semibold">{{ homestay.rating }}</span>
            <span class="text-on-surface-variant">({{ homestay.reviewCount }}条评价)</span>
          </div>
          <div class="flex items-center gap-2">
            <el-icon class="text-primary" :size="20"><Location /></el-icon>
            <span class="text-on-surface-variant">{{ homestay.location }}</span>
          </div>
        </div>

        <div class="bg-surface rounded-xl p-5 shadow-card mb-6">
          <div class="flex items-center gap-4">
            <el-avatar :size="56" class="bg-primary/10 text-primary text-lg font-bold">
              {{ homestay.host.name.charAt(0) }}
            </el-avatar>
            <div class="flex-1">
              <h3 class="font-semibold text-on-surface">房东：{{ homestay.host.name }}</h3>
              <p class="text-sm text-on-surface-variant">{{ homestay.host.badge }}</p>
            </div>
            <router-link :to="`/business/${homestay.host.id}`" class="text-primary text-sm font-medium hover:underline">
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
          <div class="space-y-4">
            <div v-for="review in homestay.reviews" :key="review.id" class="bg-surface rounded-xl p-5 shadow-card">
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
                <el-option label="1人" :value="1" />
                <el-option label="2人" :value="2" />
                <el-option label="3人" :value="3" />
                <el-option label="4人" :value="4" />
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
  </main>
</template>

<script setup>
import { ref, reactive, computed, markRaw } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { ElMessage } from 'element-plus'
import {
  StarFilled, Location, Connection, Van, IceCreamRound,
  Sunny, Coffee, Grid, House, KnifeFork
} from '@element-plus/icons-vue'

const route = useRoute()
const router = useRouter()

const bookingForm = reactive({
  checkIn: '',
  checkOut: '',
  guestCount: 2,
})

const homestay = ref({
  id: route.params.id || 1,
  title: '云顶山居·观星木屋',
  subtitle: '位于海拔1200米的云顶山，拥有绝佳观星视野和山景视野',
  category: '山景民宿',
  capacity: 4,
  rating: 4.9,
  reviewCount: 128,
  location: '浙江省·湖州市',
  price: 688,
  images: [
    'https://images.unsplash.com/photo-1600596542815-ffad4c1539a9?w=800&h=600&fit=crop',
    'https://images.unsplash.com/photo-1600585154340-be6161a56a0c?w=400&h=300&fit=crop',
    'https://images.unsplash.com/photo-1600607687939-ce8a6c25118c?w=400&h=300&fit=crop',
    'https://images.unsplash.com/photo-1600566753190-17f0baa2a6c3?w=400&h=300&fit=crop',
    'https://images.unsplash.com/photo-1600573472591-ee6b68d14c68?w=400&h=300&fit=crop',
  ],
  host: {
    id: 1,
    name: '林志远',
    badge: '已验证 · 超赞房东 · 5年经验',
  },
  description: [
    '这是一栋位于海拔1200米云顶山的独栋木屋，四周环绕着茂密的原始森林。清晨醒来，推开窗户，云海就在脚下流动；夜幕降临，满天繁星仿佛触手可及。',
    '房屋建筑面积约120平米，两室一厅一厨一卫，可入住4人。配备独立厨房、空调、地暖、24小时热水。院子里有观星平台、BBQ区域，还有一个小型的山泉水游泳池。',
  ],
  facilities: [
    { name: '高速WiFi', icon: markRaw(Connection) },
    { name: '免费停车', icon: markRaw(Van) },
    { name: '空调', icon: markRaw(IceCreamRound) },
    { name: '地暖', icon: markRaw(Sunny) },
    { name: '24小时热水', icon: markRaw(Coffee) },
    { name: '独立厨房', icon: markRaw(Grid) },
    { name: '花园庭院', icon: markRaw(House) },
    { name: '烧烤设备', icon: markRaw(KnifeFork) },
  ],
  meal: {
    name: '早餐套餐',
    price: 68,
    description: '包含：土鸡蛋、鲜牛奶、手工馒头、山区野菜、时令水果',
  },
  reviews: [
    {
      id: 1,
      author: '陈小明',
      rating: 5,
      content: '非常棒的体验！木屋很温馨，设施齐全。最惊喜的是晚上真的可以看到满天星星，房东准备的早餐也很丰富。下次还会再来！',
      date: '2024年3月入住',
    },
    {
      id: 2,
      author: '王丽华',
      rating: 5,
      content: '带家人一起过来的，老人和孩子都很喜欢。山里的空气特别好，木屋的隔音也不错。晚上一起在院子里看星星，孩子特别开心！',
      date: '2024年2月入住',
    },
  ],
})

const nights = computed(() => {
  if (!bookingForm.checkIn || !bookingForm.checkOut) return 1
  const diff = new Date(bookingForm.checkOut) - new Date(bookingForm.checkIn)
  const days = Math.ceil(diff / 86400000)
  return days > 0 ? days : 1
})

const totalPrice = computed(() => homestay.value.price * nights.value + 50 + 68)

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
.contact-house-master-btn{
  margin-left: 0;
}
</style>
