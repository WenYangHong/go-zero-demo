<template>
  <div class="min-h-screen bg-background">
    <header class="bg-surface border-b border-outline-variant/30">
      <div class="max-w-md mx-auto px-6 h-14 flex items-center justify-between">
        <router-link to="/" class="flex items-center gap-2">
          <el-icon :size="20"><ArrowLeft /></el-icon>
        </router-link>
        <span class="font-semibold text-on-surface">登录</span>
        <div class="w-5"></div>
      </div>
    </header>

    <main class="max-w-md mx-auto px-6 py-8">
      <div class="text-center mb-8">
        <div class="inline-flex items-center justify-center w-16 h-16 bg-primary/10 rounded-2xl mb-4">
          <el-icon :size="32" class="text-primary"><HomeFilled /></el-icon>
        </div>
        <h1 class="text-2xl font-bold text-on-surface">欢迎回来</h1>
        <p class="text-on-surface-variant mt-2">登录归栖民宿，探索美好旅途</p>
      </div>

      <div class="bg-surface rounded-2xl shadow-card p-6">
        <el-form ref="loginFormRef" :model="loginForm" :rules="loginRules" label-position="top">
          <el-form-item label="账号" prop="username">
            <el-input
              v-model="loginForm.username"
              placeholder="请输入手机号或邮箱"
              size="large"
              :prefix-icon="User"
            />
          </el-form-item>

          <el-form-item label="密码" prop="password">
            <el-input
              v-model="loginForm.password"
              type="password"
              placeholder="请输入密码"
              size="large"
              show-password
              :prefix-icon="Lock"
            />
          </el-form-item>

          <el-form-item label="图形验证码" prop="captcha">
            <div class="flex gap-3">
              <el-input
                v-model="loginForm.captcha"
                placeholder="请输入图形验证码"
                size="large"
                maxlength="4"
                class="flex-1"
                :prefix-icon="Key"
                @keyup.enter="handleLogin"
              />
              <div
                class="captcha-container flex items-center justify-center cursor-pointer rounded-lg overflow-hidden border border-outline-variant select-none"
                :style="{ width: '120px', height: '40px' }"
                title="点击刷新验证码"
                @click="refreshCaptcha"
              >
                <canvas ref="captchaCanvas" width="120" height="40" />
              </div>
            </div>
          </el-form-item>

          <div class="flex items-center justify-between text-sm mb-4">
            <el-checkbox v-model="loginForm.remember" label="记住我" />
            <a href="#" class="text-primary hover:underline">忘记密码？</a>
          </div>

          <el-button
            type="primary"
            size="large"
            class="w-full"
            :loading="loading"
            @click="handleLogin"
          >
            登录
          </el-button>
        </el-form>

        <div class="relative my-6">
          <div class="absolute inset-0 flex items-center">
            <div class="w-full border-t border-outline-variant/50"></div>
          </div>
          <div class="relative flex justify-center text-xs">
            <span class="px-4 bg-surface text-on-surface-variant">或</span>
          </div>
        </div>

        <div class="grid grid-cols-2 gap-3">
          <el-button size="large" @click="handleWechatLogin">
            <el-icon class="text-green-500"><ChatDotRound /></el-icon>
            <span class="ml-2">微信登录</span>
          </el-button>
          <el-button size="large" @click="handleSmsLogin">
            <el-icon class="text-blue-500"><Iphone /></el-icon>
            <span class="ml-2">短信登录</span>
          </el-button>
        </div>
      </div>

      <div class="mt-6 text-center text-sm">
        <span class="text-on-surface-variant">还没有账号？</span>
        <router-link to="/register" class="text-primary font-medium hover:underline ml-1">立即注册</router-link>
      </div>
    </main>
  </div>
</template>

<script setup>
import { ref, reactive, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { ElMessage } from 'element-plus'
import { ArrowLeft, HomeFilled, ChatDotRound, Iphone, User, Lock, Key } from '@element-plus/icons-vue'
import { useUserStore } from '@/stores/user'
import {loginByPassword} from "../api/user.js";

const router = useRouter()
const userStore = useUserStore()

const loginFormRef = ref(null)
const captchaCanvas = ref(null)
const loading = ref(false)

const loginForm = reactive({
  username: '',
  password: '',
  captcha: '',
  remember: true,
})

let captchaText = ''

const loginRules = {
  username: [
    { required: true, message: '请输入账号', trigger: 'blur' },
  ],
  password: [
    { required: true, message: '请输入密码', trigger: 'blur' },
    { min: 6, message: '密码至少6位', trigger: 'blur' },
  ],
  captcha: [
    { required: true, message: '请输入图形验证码', trigger: 'blur' },
    { len: 4, message: '图形验证码为4位', trigger: 'blur' },
  ],
}

function drawCaptcha(text) {
  const canvas = captchaCanvas.value
  if (!canvas) return
  const ctx = canvas.getContext('2d')
  const w = canvas.width
  const h = canvas.height

  ctx.fillStyle = '#f3f4f2'
  ctx.fillRect(0, 0, w, h)

  ctx.font = 'bold 24px Nunito, Noto Sans SC, sans-serif'
  ctx.textAlign = 'center'
  ctx.textBaseline = 'middle'

  const chars = text.split('')
  chars.forEach((char, i) => {
    const x = (w / 5) * (i + 1)
    const y = h / 2 + (Math.random() - 0.5) * 10
    const angle = (Math.random() - 0.5) * 0.4
    const colors = ['#059669', '#3b82f6', '#f59e0b', '#ef4444', '#8b5cf6']
    ctx.fillStyle = colors[Math.floor(Math.random() * colors.length)]
    ctx.save()
    ctx.translate(x, y)
    ctx.rotate(angle)
    ctx.fillText(char, 0, 0)
    ctx.restore()
  })

  for (let i = 0; i < 5; i++) {
    ctx.strokeStyle = `rgba(0, 0, 0, ${Math.random() * 0.1})`
    ctx.beginPath()
    ctx.moveTo(Math.random() * w, Math.random() * h)
    ctx.lineTo(Math.random() * w, Math.random() * h)
    ctx.stroke()
  }

  for (let i = 0; i < 30; i++) {
    ctx.fillStyle = `rgba(0, 0, 0, ${Math.random() * 0.15})`
    ctx.beginPath()
    ctx.arc(Math.random() * w, Math.random() * h, Math.random() * 1.5, 0, Math.PI * 2)
    ctx.fill()
  }
}

function generateCaptcha() {
  const chars = 'ABCDEFGHJKLMNPQRSTUVWXYZabcdefghjkmnpqrstuvwxyz23456789'
  let text = ''
  for (let i = 0; i < 4; i++) {
    text += chars.charAt(Math.floor(Math.random() * chars.length))
  }
  captchaText = text
  drawCaptcha(text)
}

function refreshCaptcha() {
  generateCaptcha()
}

async function handleLogin() {
  const valid = await loginFormRef.value.validate().catch(() => false)
  if (!valid) return

  if (loginForm.captcha.toLowerCase() !== captchaText.toLowerCase()) {
    ElMessage.error('图形验证码错误')
    loginForm.captcha = ''
    refreshCaptcha()
    return
  }

  loading.value = true
  try {
    const res = await loginByPassword({
      mobile: loginForm.username,
      password: loginForm.password,
    })
    userStore.setToken(res.accessToken)
    ElMessage.success('登录成功')
    router.push('/')
  } catch (e) {
    console.log("login false",e)
    refreshCaptcha()
  } finally {
    loading.value = false
  }
}

function handleWechatLogin() {
  ElMessage.info('微信登录功能开发中')
}

function handleSmsLogin() {
  ElMessage.info('短信登录功能开发中')
}

onMounted(() => {
  generateCaptcha()
})
</script>

<style scoped>
.captcha-container {
  flex-shrink: 0;
}
</style>
