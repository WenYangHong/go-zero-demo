<template>
  <main class="max-w-2xl mx-auto px-6 py-8">
    <template v-if="loading">
      <div class="bg-surface rounded-2xl shadow-card p-6 mb-6">
        <el-skeleton animated :rows="2">
          <template #template>
            <div class="flex items-center gap-3 mb-4">
              <el-skeleton-item variant="circle" style="width: 40px; height: 40px;" />
              <div class="flex-1">
                <el-skeleton-item variant="text" style="width: 50%;" />
              </div>
            </div>
            <el-skeleton-item variant="text" style="width: 80%;" />
            <el-skeleton-item variant="text" style="width: 60%; margin-top: 8px;" />
          </template>
        </el-skeleton>
      </div>

      <div class="bg-surface rounded-2xl shadow-card p-6 mb-6">
        <el-skeleton animated :rows="4">
          <template #template>
            <el-skeleton-item variant="h3" style="width: 80px; height: 24px; margin-bottom: 16px;" />
            <el-skeleton-item variant="text" style="width: 100%;" />
            <el-skeleton-item variant="text" style="width: 90%; margin-top: 12px;" />
            <el-skeleton-item variant="text" style="width: 70%; margin-top: 12px;" />
            <el-skeleton-item variant="text" style="width: 85%; margin-top: 12px;" />
          </template>
        </el-skeleton>
      </div>

      <div class="bg-surface rounded-2xl shadow-card p-6">
        <el-skeleton animated :rows="3">
          <template #template>
            <el-skeleton-item variant="h3" style="width: 80px; height: 24px; margin-bottom: 16px;" />
            <div class="flex justify-between">
              <el-skeleton-item variant="text" style="width: 40%;" />
              <el-skeleton-item variant="text" style="width: 20%;" />
            </div>
            <div class="flex justify-between" style="margin-top: 12px;">
              <el-skeleton-item variant="text" style="width: 30%;" />
              <el-skeleton-item variant="text" style="width: 25%;" />
            </div>
          </template>
        </el-skeleton>
      </div>
    </template>

    <template v-else>
      <div class="bg-surface rounded-2xl shadow-card p-6 mb-6">
        <div class="flex items-center gap-3 mb-2">
          <div class="w-10 h-10 rounded-xl bg-warning/10 flex items-center justify-center">
            <el-icon :size="20" class="text-warning"><Clock /></el-icon>
          </div>
          <div>
            <h2 class="font-bold text-on-surface text-lg">等待支付</h2>
            <p class="text-sm text-on-surface-variant">请在30分钟内完成支付，超时订单将自动取消</p>
          </div>
        </div>
        <div v-if="countdown > 0" class="mt-3 text-center">
          <span class="text-2xl font-bold text-warning font-mono">{{ countdownStr }}</span>
        </div>
        <div v-else class="mt-3 text-center">
          <span class="text-sm text-error">支付已超时，订单将自动取消</span>
        </div>
      </div>

      <div class="bg-surface rounded-2xl shadow-card p-6 mb-6">
        <h3 class="font-bold text-on-surface mb-4">订单信息</h3>
        <div class="flex gap-4 mb-4">
          <div class="w-24 h-20 rounded-lg overflow-hidden shrink-0 bg-surface-container">
            <img
              v-if="order.cover"
              :src="order.cover"
              alt="民宿图片"
              class="w-full h-full object-cover"
              @error="handleImgError"
            />
          </div>
          <div class="flex-1 min-w-0">
            <h4 class="font-semibold text-on-surface truncate">{{ order.title }}</h4>
            <p class="text-sm text-on-surface-variant mt-1">{{ order.dateRange }}</p>
            <p class="text-sm text-on-surface-variant">{{ order.guestCount }}人 · {{ order.nights }}晚</p>
          </div>
        </div>

        <div class="space-y-3 text-sm border-t border-outline-variant/30 pt-4">
          <div class="flex justify-between">
            <span class="text-on-surface-variant">订单号</span>
            <span class="text-on-surface font-mono">{{ order.sn }}</span>
          </div>
          <div class="flex justify-between">
            <span class="text-on-surface-variant">入住日期</span>
            <span class="text-on-surface">{{ order.checkIn }}</span>
          </div>
          <div class="flex justify-between">
            <span class="text-on-surface-variant">退房日期</span>
            <span class="text-on-surface">{{ order.checkOut }}</span>
          </div>
          <div class="flex justify-between">
            <span class="text-on-surface-variant">入住人数</span>
            <span class="text-on-surface">{{ order.guestCount }}人</span>
          </div>
          <div v-if="order.remark" class="flex justify-between">
            <span class="text-on-surface-variant">备注</span>
            <span class="text-on-surface">{{ order.remark }}</span>
          </div>
          <div class="flex justify-between">
            <span class="text-on-surface-variant">餐食</span>
            <span class="text-on-surface">{{ order.mealInfo }}</span>
          </div>
        </div>
      </div>

      <div class="bg-surface rounded-2xl shadow-card p-6 mb-6">
        <h3 class="font-bold text-on-surface mb-4">费用明细</h3>
        <div class="space-y-3 text-sm">
          <div class="flex justify-between">
            <span class="text-on-surface-variant">房费 ¥{{ order.price }} × {{ order.nights }}晚</span>
            <span class="text-on-surface">¥{{ order.price * order.nights }}</span>
          </div>
          <div v-if="order.mealPrice > 0" class="flex justify-between">
            <span class="text-on-surface-variant">餐食费 ¥{{ order.mealPrice }} × {{ order.guestCount }}人</span>
            <span class="text-on-surface">¥{{ order.mealPrice * order.guestCount }}</span>
          </div>
          <div class="flex justify-between">
            <span class="text-on-surface-variant">清洁费</span>
            <span class="text-on-surface">¥{{ order.cleanFee }}</span>
          </div>
          <div class="flex justify-between">
            <span class="text-on-surface-variant">服务费</span>
            <span class="text-on-surface">¥{{ order.serviceFee }}</span>
          </div>
        </div>
        <div class="flex justify-between items-center mt-4 pt-4 border-t border-outline-variant/30">
          <span class="font-semibold text-on-surface">合计</span>
          <span class="text-2xl font-bold text-primary">¥{{ order.totalAmount }}</span>
        </div>
      </div>

      <div class="bg-surface rounded-2xl shadow-card p-6 mb-6">
        <h3 class="font-bold text-on-surface mb-4">支付方式</h3>
        <el-radio-group v-model="payMethod" class="w-full">
          <div class="space-y-3 w-full">
            <label
              class="flex items-center gap-3 p-4 rounded-xl border cursor-pointer transition-all"
              :class="payMethod === 'wechat' ? 'border-primary bg-primary/5' : 'border-outline-variant hover:border-primary/50'"
            >
              <el-radio value="wechat" />
              <div class="w-8 h-8 rounded-lg bg-green-500/10 flex items-center justify-center">
                <el-icon class="text-green-500" :size="18"><ChatDotRound /></el-icon>
              </div>
              <span class="font-medium text-on-surface">微信支付</span>
            </label>
            <label
              class="flex items-center gap-3 p-4 rounded-xl border cursor-pointer transition-all"
              :class="payMethod === 'alipay' ? 'border-primary bg-primary/5' : 'border-outline-variant hover:border-primary/50'"
            >
              <el-radio value="alipay" />
              <div class="w-8 h-8 rounded-lg bg-blue-500/10 flex items-center justify-center">
                <el-icon class="text-blue-500" :size="18"><Iphone /></el-icon>
              </div>
              <span class="font-medium text-on-surface">支付宝</span>
            </label>
          </div>
        </el-radio-group>
      </div>

      <div class="flex gap-3">
        <el-button size="large" class="flex-1" @click="handleCancel">
          取消订单
        </el-button>
        <el-button
          type="primary"
          size="large"
          class="flex-1"
          :loading="payLoading"
          :disabled="countdown <= 0"
          @click="handlePay"
        >
          立即支付 ¥{{ order.totalAmount }}
        </el-button>
      </div>
    </template>
  </main>
</template>

<script setup>
import { ref, computed, onMounted, onUnmounted } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { ElMessage, ElMessageBox } from 'element-plus'
import { Clock, ChatDotRound, Iphone } from '@element-plus/icons-vue'
import { getOrderDetail, payOrder, cancelOrder } from '@/api/order.js'

const route = useRoute()
const router = useRouter()

const loading = ref(true)
const payLoading = ref(false)
const payMethod = ref('wechat')

const orderSn = route.params.sn || ''

const order = ref({
  sn: orderSn,
  title: '',
  cover: '',
  checkIn: '',
  checkOut: '',
  dateRange: '',
  nights: 1,
  guestCount: 1,
  remark: '',
  mealInfo: '不含餐',
  mealPrice: 0,
  price: 0,
  cleanFee: 50,
  serviceFee: 68,
  totalAmount: 0,
  createTime: '',
})

const countdown = ref(0)
let countdownTimer = null

const countdownStr = computed(() => {
  const m = Math.floor(countdown.value / 60)
  const s = countdown.value % 60
  return `${String(m).padStart(2, '0')}:${String(s).padStart(2, '0')}`
})

function startCountdown() {
  const deadline = new Date(order.value.createTime).getTime() + 30 * 60 * 1000
  const update = () => {
    const diff = Math.floor((deadline - Date.now()) / 1000)
    countdown.value = diff > 0 ? diff : 0
  }
  update()
  countdownTimer = setInterval(update, 1000)
}

function formatDate(ts) {
  if (!ts) return ''
  const d = new Date(typeof ts === 'number' ? ts * 1000 : ts)
  const y = d.getFullYear()
  const m = String(d.getMonth() + 1).padStart(2, '0')
  const day = String(d.getDate()).padStart(2, '0')
  return `${y}-${m}-${day}`
}

async function fetchOrderDetail() {
  try {
    const res = await getOrderDetail(orderSn)
    const d = res.order || res
    const checkIn = formatDate(d.live_start_time)
    const checkOut = formatDate(d.live_end_time)
    const nights = d.live_end_time && d.live_start_time
      ? Math.ceil((new Date(d.live_end_time * 1000) - new Date(d.live_start_time * 1000)) / 86400000)
      : 1

    order.value = {
      sn: d.sn || orderSn,
      title: d.title || d.homestay_title || '民宿预订',
      cover: d.cover || d.img_info || '',
      checkIn,
      checkOut,
      dateRange: `${checkIn} 至 ${checkOut}`,
      nights,
      guestCount: d.people_num || 1,
      remark: d.remark || '',
      mealInfo: d.is_food ? (d.food_info || '含早餐') : '不含餐',
      mealPrice: d.food_price || 0,
      price: d.homestay_price || 0,
      cleanFee: 50,
      serviceFee: 68,
      totalAmount: d.total_price || d.order_price || 0,
      createTime: d.create_time || new Date().toISOString(),
    }

    startCountdown()
  } catch {
    order.value = {
      ...order.value,
      title: '民宿预订',
      totalAmount: 0,
    }
  }
}

async function handlePay() {
  try {
    await ElMessageBox.confirm(
      `<div style="text-align: center; padding: 8px 0;">
        <div style="font-size: 16px; margin-bottom: 8px;">确认支付</div>
        <div style="font-size: 28px; font-weight: 700; color: #059669;">¥${order.value.totalAmount}</div>
        <div style="color: #6B7280; margin-top: 4px;">${payMethod.value === 'wechat' ? '微信支付' : '支付宝'}</div>
      </div>`,
      '支付确认',
      {
        confirmButtonText: '确认支付',
        cancelButtonText: '取消',
        dangerouslyUseHTMLString: true,
      }
    )
  } catch {
    return
  }

  payLoading.value = true
  try {
    await payOrder({
      sn: order.value.sn,
      pay_method: payMethod.value,
    })

    ElMessageBox.alert(
      '<div style="text-align: center; padding: 16px 0;"><div style="font-size: 48px; margin-bottom: 12px;">✅</div><div style="font-size: 18px; font-weight: 600; margin-bottom: 8px;">支付成功！</div><div style="color: #6B7280;">房东将尽快确认您的订单，祝您旅途愉快</div></div>',
      '支付结果',
      {
        confirmButtonText: '返回首页',
        dangerouslyUseHTMLString: true,
        callback: () => {
          router.push('/')
        },
      }
    )
  } catch (err) {
    ElMessageBox.alert(
      `<div style="text-align: center; padding: 16px 0;"><div style="font-size: 48px; margin-bottom: 12px;">😔</div><div style="font-size: 18px; font-weight: 600; margin-bottom: 8px;">支付失败</div><div style="color: #6B7280;">${err?.message || '支付遇到问题，请稍后重试'}</div></div>`,
      '支付结果',
      {
        confirmButtonText: '重试',
        dangerouslyUseHTMLString: true,
      }
    )
  } finally {
    payLoading.value = false
  }
}

async function handleCancel() {
  try {
    await ElMessageBox.confirm(
      '确定要取消此订单吗？取消后需重新预订。',
      '取消订单',
      {
        confirmButtonText: '确认取消',
        cancelButtonText: '再想想',
        type: 'warning',
      }
    )
  } catch {
    return
  }

  try {
    await cancelOrder(order.value.sn)
    ElMessage.success('订单已取消')
    router.push('/')
  } catch (err) {
    ElMessage.error(err?.message || '取消失败，请稍后重试')
  }
}

function handleImgError(e) {
  e.target.src = 'data:image/svg+xml,' + encodeURIComponent(
    `<svg xmlns="http://www.w3.org/2000/svg" width="96" height="80"><rect width="100%" height="100%" fill="#f3f4f6"/><text x="50%" y="50%" text-anchor="middle" font-size="16" fill="#9ca3af" dominant-baseline="central">🖼</text></svg>`
  )
}

onMounted(async () => {
  await fetchOrderDetail()
  loading.value = false
})

onUnmounted(() => {
  if (countdownTimer) {
    clearInterval(countdownTimer)
  }
})
</script>
