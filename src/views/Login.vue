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
          <el-form-item label="手机号" prop="phone">
            <el-input
              v-model="loginForm.phone"
              placeholder="请输入手机号"
              size="large"
              maxlength="11"
            >
              <template #append>
                <el-button :disabled="codeCooldown > 0" @click="handleSendCode">
                  {{ codeCooldown > 0 ? `${codeCooldown}秒后重试` : '获取验证码' }}
                </el-button>
              </template>
            </el-input>
          </el-form-item>

          <el-form-item label="验证码" prop="code">
            <el-input
              v-model="loginForm.code"
              placeholder="请输入验证码"
              size="large"
              maxlength="6"
            />
          </el-form-item>

          <div class="flex items-center justify-between text-sm mb-4">
            <el-checkbox v-model="loginForm.remember" label="记住我" />
            <a href="#" class="text-primary hover:underline">忘记密码？</a>
          </div>

          <el-button type="primary" size="large" class="w-full" :loading="loading" @click="handleLogin">
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
import { ref, reactive } from 'vue'
import { useRouter } from 'vue-router'
import { ElMessage } from 'element-plus'
import { ArrowLeft, HomeFilled, ChatDotRound, Iphone } from '@element-plus/icons-vue'
import { useUserStore } from '@/stores/user'

const router = useRouter()
const userStore = useUserStore()

const loginFormRef = ref(null)
const loading = ref(false)
const codeCooldown = ref(0)

const loginForm = reactive({
  phone: '',
  code: '',
  remember: true,
})

const loginRules = {
  phone: [
    { required: true, message: '请输入手机号', trigger: 'blur' },
    { pattern: /^1[3-9]\d{9}$/, message: '请输入正确的手机号', trigger: 'blur' },
  ],
  code: [
    { required: true, message: '请输入验证码', trigger: 'blur' },
  ],
}

function handleSendCode() {
  if (!loginForm.phone) {
    ElMessage.warning('请先输入手机号')
    return
  }
  if (!/^1[3-9]\d{9}$/.test(loginForm.phone)) {
    ElMessage.warning('请输入正确的手机号')
    return
  }
  codeCooldown.value = 60
  const timer = setInterval(() => {
    codeCooldown.value--
    if (codeCooldown.value <= 0) {
      clearInterval(timer)
    }
  }, 1000)
  ElMessage.success('验证码已发送')
}

async function handleLogin() {
  const valid = await loginFormRef.value.validate().catch(() => false)
  if (!valid) return

  loading.value = true
  try {
    // TODO: 调用后端登录接口
    // const res = await loginByCode({ phone: loginForm.phone, code: loginForm.code })
    // userStore.setToken(res.token)
    ElMessage.success('登录成功')
    router.push('/')
  } catch {
    // 错误已在拦截器中处理
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
</script>
