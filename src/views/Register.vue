<template>
  <div class="min-h-screen bg-background">
    <header class="bg-surface border-b border-outline-variant/30">
      <div class="max-w-md mx-auto px-6 h-14 flex items-center justify-between">
        <router-link to="/" class="flex items-center gap-2">
          <el-icon :size="20"><ArrowLeft /></el-icon>
        </router-link>
        <span class="font-semibold text-on-surface">注册</span>
        <div class="w-5"></div>
      </div>
    </header>

    <main class="max-w-md mx-auto px-6 py-8">
      <div class="text-center mb-8">
        <div class="inline-flex items-center justify-center w-16 h-16 bg-primary/10 rounded-2xl mb-4">
          <el-icon :size="32" class="text-primary"><HomeFilled /></el-icon>
        </div>
        <h1 class="text-2xl font-bold text-on-surface">创建账号</h1>
        <p class="text-on-surface-variant mt-2">加入归栖，开启美好旅途</p>
      </div>

      <div class="bg-surface rounded-2xl shadow-card p-6">
        <el-form ref="registerFormRef" :model="registerForm" :rules="registerRules" label-position="top">
          <el-form-item label="手机号" prop="phone">
            <el-input
              v-model="registerForm.phone"
              placeholder="请输入手机号"
              size="large"
              maxlength="11"
            />
          </el-form-item>

          <el-form-item label="验证码" prop="code">
            <el-input
              v-model="registerForm.code"
              placeholder="请输入验证码"
              size="large"
              maxlength="6"
            >
              <template #append>
                <el-button :disabled="codeCooldown > 0" @click="handleSendCode">
                  {{ codeCooldown > 0 ? `${codeCooldown}秒后重试` : '获取验证码' }}
                </el-button>
              </template>
            </el-input>
          </el-form-item>
          <el-form-item label="昵称" prop="nickname">
            <el-input
                v-model="registerForm.nickname"
                type="text"
                placeholder="请设置账户昵称"
                size="large"
            />
          </el-form-item>
          <el-form-item label="设置密码" prop="password">
            <el-input
              v-model="registerForm.password"
              type="password"
              placeholder="请设置密码（6位以上）"
              size="large"
              show-password
            />
          </el-form-item>

          <el-form-item label="确认密码" prop="confirmPassword">
            <el-input
              v-model="registerForm.confirmPassword"
              type="password"
              placeholder="请再次输入密码"
              size="large"
              show-password
            />
          </el-form-item>

          <el-form-item>
            <el-checkbox v-model="registerForm.agree">
              <span class="text-xs text-on-surface-variant leading-relaxed">
                我已阅读并同意
                <a href="#" class="text-primary">《用户服务协议》</a>
                和
                <a href="#" class="text-primary">《隐私政策》</a>
              </span>
            </el-checkbox>
          </el-form-item>

          <el-button type="primary" size="large" class="w-full" :loading="loading" @click="handleRegister">
            注册
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

        <el-button size="large" class="w-full" @click="handleWechatRegister">
          <el-icon class="text-green-500"><ChatDotRound /></el-icon>
          <span class="ml-2">微信一键注册</span>
        </el-button>
      </div>

      <div class="mt-6 text-center text-sm">
        <span class="text-on-surface-variant">已有账号？</span>
        <router-link to="/login" class="text-primary font-medium hover:underline ml-1">立即登录</router-link>
      </div>
    </main>
  </div>
</template>

<script setup>
import { ref, reactive } from 'vue'
import { useRouter } from 'vue-router'
import { ElMessage } from 'element-plus'
import { ArrowLeft, HomeFilled, ChatDotRound } from '@element-plus/icons-vue'
import {register} from "../api/user.js";
import { useUserStore } from '@/stores/user'
const router = useRouter()
const userStore = useUserStore()
const registerFormRef = ref(null)
const loading = ref(false)
const codeCooldown = ref(0)

const registerForm = reactive({
  phone: '',
  code: '',
  password: '',
  nickname: '',
  confirmPassword: '',
  agree: false,
})

const validateConfirmPassword = (rule, value, callback) => {
  if (value !== registerForm.password) {
    callback(new Error('两次输入的密码不一致'))
  } else {
    callback()
  }
}

const registerRules = {
  phone: [
    { required: true, message: '请输入手机号', trigger: 'blur' },
    { pattern: /^1[3-9]\d{9}$/, message: '请输入正确的手机号', trigger: 'blur' },
  ],
  code: [
    { required: true, message: '请输入验证码', trigger: 'blur' },
  ],
  nickname: [
    { required: true, message: '请输入昵称', trigger: 'blur' },
  ],
  password: [
    { required: true, message: '请设置密码', trigger: 'blur' },
    { min: 6, message: '密码至少6位', trigger: 'blur' },
  ],
  confirmPassword: [
    { required: true, message: '请确认密码', trigger: 'blur' },
    { validator: validateConfirmPassword, trigger: 'blur' },
  ],
}

function handleSendCode() {
  if (!registerForm.phone) {
    ElMessage.warning('请先输入手机号')
    return
  }
  if (!/^1[3-9]\d{9}$/.test(registerForm.phone)) {
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

async function handleRegister() {
  const valid = await registerFormRef.value.validate().catch(() => false)
  if (!valid) return

  if (!registerForm.agree) {
    ElMessage.warning('请阅读并同意用户协议和隐私政策')
    return
  }

  loading.value = true
  try {
    // TODO: 调用后端注册接口
    const res = await register({ mobile: registerForm.phone, nick_name: registerForm.nickname, password: registerForm.password })
    ElMessage.success('注册成功')
    userStore.setToken(res.accessToken)
    router.push('/')
  } catch (e){
    // 错误已在拦截器中处理
    console.error(e)
  } finally {
    loading.value = false
  }
}

function handleWechatRegister() {
  ElMessage.info('微信注册功能开发中')
}
</script>
