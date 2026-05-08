<template>
  <main class="max-w-7xl mx-auto px-6 py-8">
    <template v-if="loading">
      <div class="mb-8">
        <el-skeleton animated>
          <template #template>
            <el-skeleton-item variant="rect" style="width: 100%; height: 256px; border-radius: 1rem;" />
          </template>
        </el-skeleton>
      </div>

      <div class="grid grid-cols-1 md:grid-cols-3 gap-8">
        <div class="md:col-span-2">
          <el-skeleton animated :rows="3">
            <template #template>
              <div class="bg-surface rounded-xl p-6 shadow-card mb-6">
                <div class="flex items-start gap-4 mb-4">
                  <el-skeleton-item variant="circle" style="width: 64px; height: 64px;" />
                  <div class="flex-1">
                    <el-skeleton-item variant="text" style="width: 30%;" />
                    <el-skeleton-item variant="text" style="width: 50%; margin-top: 8px;" />
                  </div>
                </div>
                <el-skeleton-item variant="text" style="width: 100%;" />
                <el-skeleton-item variant="text" style="width: 90%; margin-top: 8px;" />
              </div>

              <div class="bg-surface rounded-xl p-6 shadow-card mb-6">
                <el-skeleton-item variant="h3" style="width: 80px; height: 24px; margin-bottom: 16px;" />
                <el-skeleton-item variant="text" style="width: 100%;" />
                <el-skeleton-item variant="text" style="width: 60%; margin-top: 8px;" />
                <el-skeleton-item variant="text" style="width: 80%; margin-top: 8px;" />
                <el-skeleton-item variant="text" style="width: 50%; margin-top: 8px;" />
              </div>

              <div class="bg-surface rounded-xl p-6 shadow-card mb-6">
                <el-skeleton-item variant="h3" style="width: 80px; height: 24px; margin-bottom: 16px;" />
                <div class="grid grid-cols-2 md:grid-cols-4 gap-4">
                  <el-skeleton-item v-for="i in 4" :key="i" variant="text" style="width: 100%;" />
                </div>
              </div>
            </template>
          </el-skeleton>
        </div>

        <div>
          <el-skeleton animated :rows="3">
            <template #template>
              <el-skeleton-item variant="h3" style="width: 80px; height: 24px; margin-bottom: 16px;" />
              <div v-for="i in 3" :key="i" style="margin-bottom: 16px;">
                <el-skeleton-item variant="rect" style="width: 100%; height: 144px; border-radius: 0.75rem; margin-bottom: 8px;" />
                <el-skeleton-item variant="text" style="width: 60%;" />
                <el-skeleton-item variant="text" style="width: 40%; margin-top: 4px;" />
              </div>
            </template>
          </el-skeleton>
        </div>
      </div>
    </template>

    <template v-else>
      <div class="relative mb-8">
        <img
          :src="business.cover"
          alt="店铺封面"
          class="w-full h-64 object-cover rounded-2xl"
          @error="handleImgError"
        />
        <div class="absolute inset-0 bg-gradient-to-t from-black/50 to-transparent rounded-2xl"></div>
        <div class="absolute bottom-6 left-6 right-6">
          <div class="flex items-end gap-4">
            <div class="w-24 h-24 rounded-xl overflow-hidden border-4 border-white shadow-lg">
              <img
                :src="business.avatar"
                alt="房东头像"
                class="w-full h-full object-cover"
                @error="handleImgError"
              />
            </div>
            <div class="flex-1 text-white mb-1">
              <h1 class="text-2xl font-bold mb-1">{{ business.title }}</h1>
              <p class="text-white/80 text-sm">{{ business.badge }}</p>
            </div>
          </div>
        </div>
      </div>

      <div class="grid grid-cols-1 md:grid-cols-3 gap-8">
        <div class="md:col-span-2">
          <section class="bg-surface rounded-xl p-6 shadow-card mb-6">
            <h2 class="text-lg font-bold text-on-surface mb-4">房东介绍</h2>
            <div class="flex items-start gap-4 mb-4">
              <el-avatar :size="64" class="bg-primary/10 text-primary text-xl font-bold">
                {{ business.hostName.charAt(0) }}
              </el-avatar>
              <div>
                <h3 class="font-semibold text-on-surface text-lg">{{ business.hostName }}</h3>
                <p class="text-sm text-on-surface-variant">{{ business.hostJoinInfo }}</p>
                <div class="flex items-center gap-4 mt-2">
                  <div class="flex items-center gap-1">
                    <el-icon class="text-warning"><StarFilled /></el-icon>
                    <span class="text-sm font-medium">{{ business.rating }}</span>
                    <span class="text-sm text-on-surface-variant">评分</span>
                  </div>
                  <div class="flex items-center gap-1">
                    <el-icon class="text-primary"><ChatDotRound /></el-icon>
                    <span class="text-sm font-medium">{{ business.replyRate }}</span>
                    <span class="text-sm text-on-surface-variant">回复率</span>
                  </div>
                </div>
              </div>
            </div>
            <p class="text-on-surface-variant leading-relaxed">{{ business.hostDescription }}</p>
          </section>

          <section class="bg-surface rounded-xl p-6 shadow-card mb-6">
            <h2 class="text-lg font-bold text-on-surface mb-4">店铺介绍</h2>
            <p class="text-on-surface-variant leading-relaxed mb-4">{{ business.shopIntro }}</p>
            <ul class="space-y-3 text-on-surface-variant">
              <li v-for="item in business.homestayList" :key="item.name" class="flex items-start gap-3">
                <el-icon class="text-primary shrink-0 mt-0.5" :size="20"><CircleCheckFilled /></el-icon>
                <span><strong>{{ item.name }}</strong> - {{ item.description }}</span>
              </li>
            </ul>
          </section>

          <section class="bg-surface rounded-xl p-6 shadow-card mb-6">
            <h2 class="text-lg font-bold text-on-surface mb-4">认证信息</h2>
            <div class="grid grid-cols-2 md:grid-cols-4 gap-4">
              <div v-for="cert in business.certifications" :key="cert.name" class="flex items-center gap-2 text-sm">
                <el-icon :size="20" :class="cert.color"><component :is="cert.icon" /></el-icon>
                <span class="text-on-surface-variant">{{ cert.name }}</span>
              </div>
            </div>
          </section>

          <section>
            <div class="flex items-center justify-between mb-4">
              <h2 class="text-lg font-bold text-on-surface">评价概览</h2>
              <router-link :to="`/comment/${business.id}?type=homestayBusiness`" class="text-primary text-sm font-medium hover:underline">
                查看全部
              </router-link>
            </div>

            <div class="grid grid-cols-2 md:grid-cols-4 gap-4 mb-6">
              <div v-for="score in business.scoreOverview" :key="score.label" class="bg-surface rounded-xl p-4 text-center shadow-card">
                <div class="flex items-center justify-center gap-1 mb-2">
                  <el-icon :size="16" :class="score.iconColor"><component :is="score.icon" /></el-icon>
                  <span class="text-xl font-bold text-on-surface">{{ score.value }}</span>
                </div>
                <p class="text-xs text-on-surface-variant">{{ score.label }}</p>
              </div>
            </div>

            <div v-if="business.reviews.length > 0" class="space-y-4">
              <div v-for="review in business.reviews" :key="review.id" class="bg-surface rounded-xl p-5 shadow-card">
                <div class="flex items-start gap-4">
                  <el-avatar :size="40" class="bg-primary/10 text-primary font-bold shrink-0">
                    {{ review.author.charAt(0) }}
                  </el-avatar>
                  <div class="flex-1">
                    <div class="flex items-center gap-2 mb-2">
                      <span class="font-medium">{{ review.author }}</span>
                      <span class="text-xs text-on-surface-variant">{{ review.date }}</span>
                    </div>
                    <p class="text-sm text-on-surface-variant">{{ review.content }}</p>
                  </div>
                </div>
              </div>
            </div>
            <div v-else class="bg-surface rounded-xl p-10 shadow-card text-center text-on-surface-variant">
              暂无评价
            </div>
          </section>
        </div>

        <div>
          <h3 class="font-bold text-on-surface mb-4">Ta的民宿</h3>
          <div v-if="business.homestayList.length > 0" class="space-y-4">
            <router-link
              v-for="item in business.homestayList"
              :key="item.id"
              :to="`/homestay/${item.id}`"
              class="block bg-surface rounded-xl overflow-hidden shadow-card hover:shadow-float transition-all"
            >
              <img
                :src="item.cover"
                :alt="item.name"
                class="w-full h-36 object-cover"
                @error="handleImgError"
              />
              <div class="p-4">
                <h4 class="font-semibold text-on-surface mb-1">{{ item.name }}</h4>
                <div class="flex items-center gap-2 mb-2">
                  <span class="px-2 py-0.5 bg-primary/10 text-primary text-xs rounded">{{ item.category }}</span>
                  <span class="text-xs text-on-surface-variant">可住{{ item.capacity }}人</span>
                </div>
                <div class="flex items-center justify-between">
                  <div class="flex items-center gap-1">
                    <el-icon class="text-warning" :size="14"><StarFilled /></el-icon>
                    <span class="text-sm">{{ item.rating }}</span>
                  </div>
                  <span class="text-primary font-bold">¥{{ item.price }}<span class="text-xs text-on-surface-variant">/晚</span></span>
                </div>
              </div>
            </router-link>
          </div>
          <div v-else class="text-center text-on-surface-variant py-8">
            暂无民宿
          </div>
        </div>
      </div>
    </template>
  </main>
</template>

<script setup>
import { ref, markRaw, onMounted } from 'vue'
import { useRoute } from 'vue-router'
import {
  StarFilled, ChatDotRound, CircleCheckFilled,
  SuccessFilled, Medal, Trophy, Stamp
} from '@element-plus/icons-vue'
import { getHomestayBossWithId } from '@/api/homestay.js'

const route = useRoute()

const loading = ref(true)

function trim(s) {
  return (s || '').replace(/^\n+|\n+$/g, '').trim()
}

function mapBossData(data) {
  const star = data.star || 5
  return {
    id: data.id,
    cover: data.cover,
    avatar: data.HeaderImg,
    title: trim(data.title),
    badge: '已验证 · 超赞房东',
    hostName: trim(data.title),
    hostJoinInfo: '2019年加入平台 · 已帮助500+位房客',
    rating: star,
    replyRate: '100%',
    hostDescription: trim(data.boss_info),
    shopIntro: trim(data.info),
    certifications: [
      { name: '身份证已认证', icon: markRaw(SuccessFilled), color: 'text-success' },
      { name: '营业执照', icon: markRaw(CircleCheckFilled), color: 'text-success' },
      { name: '超赞房东', icon: markRaw(Trophy), color: 'text-warning' },
      { name: '五良民宿', icon: markRaw(Stamp), color: 'text-primary' },
    ],
    scoreOverview: [
      { label: '整体评分', value: star, icon: markRaw(StarFilled), iconColor: 'text-warning' },
      { label: '清洁程度', value: 4.9, icon: markRaw(Medal), iconColor: 'text-primary' },
      { label: '位置交通', value: 4.8, icon: markRaw(StarFilled), iconColor: 'text-primary' },
      { label: '房东服务', value: 5.0, icon: markRaw(Medal), iconColor: 'text-primary' },
    ],
    reviews: [
      {
        id: 1,
        author: '陈小明',
        date: '2024-03-15',
        content: '房东人非常热情，提前联系我们告知路线和天气情况。民宿比照片上还要漂亮，非常干净整洁。下次来浙江还会选择这里！',
      },
    ],
    homestayList: [
      {
        id: 1,
        name: '云顶山居·观星木屋',
        cover: 'https://images.unsplash.com/photo-1600596542815-ffad4c1539a9?w=400&h=250&fit=crop',
        category: '山景',
        capacity: 4,
        rating: 4.9,
        price: 688,
        description: '海拔1200米，两室一厅，可住4人，适合家庭出游',
      },
      {
        id: 4,
        name: '林间小屋·森林树屋',
        cover: 'https://images.unsplash.com/photo-1518780664697-55e3ad937233?w=400&h=250&fit=crop',
        category: '森林',
        capacity: 2,
        rating: 5.0,
        price: 588,
        description: '建造在古树上，一室一厅，可住2人，适合情侣',
      },
      {
        id: 7,
        name: '山泉雅居·私汤民宿',
        cover: 'https://images.unsplash.com/photo-1540518614846-7eded433c457?w=400&h=250&fit=crop',
        category: '温泉',
        capacity: 6,
        rating: 4.8,
        price: 1288,
        description: '带私人温泉，三室一厅，可住6人，适合团建',
      },
    ],
  }
}

const business = ref(mapBossData({}))

async function fetchBusinessDetail() {
  try {
    const res = await getHomestayBossWithId({ id: route.params.id })
    business.value = mapBossData(res)
  } catch {
    business.value = mapBossData({})
  }
}

function handleImgError(e) {
  e.target.src = 'data:image/svg+xml,' + encodeURIComponent(
    `<svg xmlns="http://www.w3.org/2000/svg" width="400" height="300"><rect width="100%" height="100%" fill="#f3f4f6"/><text x="50%" y="50%" text-anchor="middle" font-size="24" fill="#9ca3af" dominant-baseline="central">🖼</text></svg>`
  )
}

onMounted(async () => {
  await fetchBusinessDetail()
  loading.value = false
})
</script>
