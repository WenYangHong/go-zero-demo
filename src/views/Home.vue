<template>
  <main class="max-w-7xl mx-auto px-6 py-8">
    <div class="bg-surface rounded-2xl shadow-card p-6 mb-8">
      <div class="flex flex-col md:flex-row gap-4">
        <div class="flex-1">
          <label class="block text-sm font-medium text-on-surface mb-2">目的地</label>
          <el-input
            v-model="searchForm.destination"
            placeholder="想去哪里？"
            size="large"
            class="w-full"
          />
        </div>
        <div class="w-full md:w-56">
          <label class="block text-sm font-medium text-on-surface mb-2">入住日期</label>
          <el-date-picker
            v-model="searchForm.checkIn"
            type="date"
            placeholder="选择入住日期"
            size="large"
            class="w-full"
            :disabled-date="disablePastDate"
            value-format="YYYY-MM-DD"
          />
        </div>
        <div class="w-full md:w-56">
          <label class="block text-sm font-medium text-on-surface mb-2">退房日期</label>
          <el-date-picker
            v-model="searchForm.checkOut"
            type="date"
            placeholder="选择退房日期"
            size="large"
            class="w-full"
            :disabled-date="disableCheckOutDate"
            value-format="YYYY-MM-DD"
          />
        </div>
        <div class="w-full md:w-36">
          <label class="block text-sm font-medium text-on-surface mb-2">人数</label>
          <el-select v-model="searchForm.guestCount" size="large" class="w-full">
            <el-option label="1人" :value="1" />
            <el-option label="2人" :value="2" />
            <el-option label="3人" :value="3" />
            <el-option label="4人以上" :value="4" />
          </el-select>
        </div>
        <div class="flex items-end">
          <el-button type="primary" size="large" class="w-full md:w-auto px-8" @click="handleSearch">
            搜索民宿
          </el-button>
        </div>
      </div>
    </div>

    <section class="mb-10">
      <h2 class="text-xl font-bold text-on-surface mb-4">热门分类</h2>
      <div class="flex gap-4 flex-wrap">
        <a
          v-for="cat in categories"
          :key="cat.name"
          href="#"
          class="inline-flex items-center gap-2 px-5 py-2.5 bg-surface rounded-full shadow-card hover:shadow-float transition-shadow"
          @click.prevent="handleCategoryClick(cat.name)"
        >
          <el-icon class="text-primary"><component :is="cat.icon" /></el-icon>
          <span class="text-sm font-medium">{{ cat.name }}</span>
        </a>
      </div>
    </section>

    <section>
      <div class="flex items-center justify-between mb-6">
        <h2 class="text-xl font-bold text-on-surface">精选民宿</h2>
        <a href="#" class="text-sm text-primary font-medium hover:underline">查看更多</a>
      </div>
      <div v-loading="listLoading" class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6 min-h-[200px]">
        <HomestayCard
          v-for="item in homestayList"
          :key="item.id"
          :homestay="item"
        />
      </div>
      <div v-if="!listLoading && homestayList.length === 0" class="text-center py-16 text-on-surface-variant">
        暂无民宿数据
      </div>
    </section>
  </main>
</template>

<script setup>
import { ref, reactive, markRaw, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { Sunny, Coffee, Grid, House, Van } from '@element-plus/icons-vue'
import HomestayCard from '@/components/HomestayCard.vue'
import { getHomestayList } from '@/api/homestay'
const router = useRouter()

const searchForm = reactive({
  destination: '',
  checkIn: '',
  checkOut: '',
  guestCount: 2,
})

const categories = [
  { name: '山景民宿', icon: markRaw(Sunny) },
  { name: '海景民宿', icon: markRaw(Coffee) },
  { name: '森林民宿', icon: markRaw(Grid) },
  { name: '古镇民宿', icon: markRaw(House) },
  { name: '露营民宿', icon: markRaw(Van) },
]

const homestayList = ref([])
const listLoading = ref(false)

// const homestayList = ref([
//   {
//     id: 1,
//     title: '云顶山居·观星木屋',
//     cover: 'https://images.unsplash.com/photo-1600596542815-ffad4c1539a9?w=800&h=600&fit=crop',
//     tag: '精选',
//     category: '山景',
//     capacity: 4,
//     description: '位于海拔1200米的云顶山居，远离城市喧嚣，拥有绝佳观星视野。',
//     rating: 4.9,
//     reviewCount: 128,
//     price: 688,
//   },
//   {
//     id: 2,
//     title: '蓝湾海岸·度假别墅',
//     cover: 'https://images.unsplash.com/photo-1566073771259-6a8506099945?w=800&h=600&fit=crop',
//     tag: '海景',
//     category: '海景',
//     capacity: 6,
//     description: '直面大海的度假别墅，私人沙滩，泳池别墅，适合家庭出游。',
//     rating: 4.8,
//     reviewCount: 96,
//     price: 1288,
//   },
//   {
//     id: 3,
//     title: '江南水乡·临河小院',
//     cover: 'https://images.unsplash.com/photo-1587061949409-02df41d5e562?w=800&h=600&fit=crop',
//     tag: '古镇',
//     category: '古镇',
//     capacity: 2,
//     description: '位于乌镇景区内，临河而建，古色古香，感受江南水乡风情。',
//     rating: 4.7,
//     reviewCount: 86,
//     price: 458,
//   },
//   {
//     id: 4,
//     title: '绿野仙踪·森林树屋',
//     cover: 'https://images.unsplash.com/photo-1518780664697-55e3ad937233?w=800&h=600&fit=crop',
//     tag: '森林',
//     category: '森林',
//     capacity: 3,
//     description: '建造在百年古树上的树屋，与自然融为一体，聆听鸟语花香。',
//     rating: 4.9,
//     reviewCount: 156,
//     price: 888,
//   },
//   {
//     id: 5,
//     title: '禅意居·日式庭院',
//     cover: 'https://images.unsplash.com/photo-1499696010180-025ef6e1a8f9?w=800&h=600&fit=crop',
//     tag: '精品',
//     category: '日式',
//     capacity: 4,
//     description: '日式枯山水庭院，私汤温泉，体验日式禅意生活方式。',
//     rating: 4.8,
//     reviewCount: 203,
//     price: 768,
//   },
//   {
//     id: 6,
//     title: '星空露营·帐篷营地',
//     cover: 'https://images.unsplash.com/photo-1520250497591-112f2f40a3f4?w=800&h=600&fit=crop',
//     tag: '露营',
//     category: '露营',
//     capacity: 2,
//     description: '远离光污染的高原营地，夜晚银河清晰可见，浪漫至极。',
//     rating: 4.6,
//     reviewCount: 72,
//     price: 328,
//   },
// ])

async function fetchList() {
  listLoading.value = true
  try {
    const res = await getHomestayList({
      "page":1,
      "pageSize": 6,
    })
    homestayList.value = res.list || res || []
  } catch {
    homestayList.value = []
  } finally {
    listLoading.value = false
  }
}

onMounted(() => {
  fetchList()
})

function disablePastDate(date) {
  return date.getTime() < Date.now() - 86400000
}

function disableCheckOutDate(date) {
  if (!searchForm.checkIn) return date.getTime() < Date.now() - 86400000
  return date.getTime() <= new Date(searchForm.checkIn).getTime()
}

function handleSearch() {
  router.push({
    path: '/',
    query: {
      destination: searchForm.destination || undefined,
      checkIn: searchForm.checkIn || undefined,
      checkOut: searchForm.checkOut || undefined,
      guests: searchForm.guestCount,
    },
  })
}

function handleCategoryClick(name) {
  router.push({ path: '/', query: { category: name } })
}
</script>
