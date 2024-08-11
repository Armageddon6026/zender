<script setup lang="ts">
import { reactive, ref, onMounted } from 'vue'
import router from '@/router'
import { useUserStore } from '@/stores/UserStore'
import { useServicesStore } from '@/stores/ServicesStore'
import { useEventStore } from '@/stores/EventStore'
import { ClearAllCookies } from '@/common/cookie'

interface loginForm {
  Account: string
  Password: string
}

const ourUserStore = useUserStore()
const ourServicesStore = useServicesStore()
const ourEventStore = useEventStore()
const ourLoginInput: loginForm = reactive({ Account: '', Password: '' })
const ourErrorMessage = ref('')

onMounted(() => {
  ClearAllCookies()
  ourUserStore.$reset()
  ourServicesStore.$reset()
  ourEventStore.$reset()
})

async function handleLogin() {
  try {
    await ourUserStore.Login(ourLoginInput.Account, ourLoginInput.Password)
    router.push({ name: 'HomeView' })
  } catch (err) {
    ourErrorMessage.value = String(err)
  }
}
</script>

<template>
  <div class="LoginView-container">
    <div class="login-box">
      <div class="input-form">
        <h2 c="white" text="center" font="bold size-34px" m="0">Login</h2>
        <div class="input-box">
          <input
            v-model="ourLoginInput.Account"
            type="text"
            placeholder="account"
            data-space-bottom="0.5rem"
            @keyup.enter="handleLogin"
          />
        </div>
        <div class="input-box">
          <input
            v-model="ourLoginInput.Password"
            type="password"
            placeholder="Password"
            data-space-bottom="0.5rem"
            @keyup.enter="handleLogin"
          />
        </div>
        <div float-right class="remember-checkbox"><a href="">Create Zender ID ></a></div>
        <div>
          <button class="submit-button" @click="handleLogin">Login</button>
          <button
            bg="accept-dark"
            c="white"
            mt="10px"
            un-border="solid 2px accept-normal rd-4px"
            w="100%"
            h="45px"
            font="size-20px 600"
            hover="bg-accept-normal"
          >
            SSO
          </button>
        </div>
        <div class="message-font">
          <span>{{ ourErrorMessage }} </span>
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>
.LoginView-container {
  display: flex;
  justify-content: center;
  align-items: center;
  min-height: 80vh;
}

.LoginView-container .login-box {
  background-color: rgba(255, 255, 255, 0.1);
  position: relative;
  width: 400px;
  height: 450px;
  border-radius: 10px;
  display: flex;
  justify-content: center;
  align-items: center;
}

.LoginView-container .input-box {
  position: relative;
  margin: 30px 0;
  width: 310px;
  border-bottom: 2px solid white;
}

.LoginView-container .input-box input {
  width: 100%;
  height: 50px;
  background-color: transparent;
  border: none;
  outline: none;
  font-size: 20px;
  padding: 0 35px 0 5px;
  color: var(--main-color);
}

.LoginView-container .remember-checkbox {
  margin: -15px 0 15px;
  justify-content: space-between;
}

.LoginView-container .remember-checkbox a {
  text-decoration: none;
  color: var(--main-color);
  font-size: 14px;
}

.LoginView-container .message-font {
  color: var(--main-deny-color);
  text-align: center;
  margin-top: 5px;
  letter-spacing: 3px;
  width: 300px;
}

.LoginView-container .submit-button {
  background-color: var(--main-color-dark);
  border: 2px solid var(--main-color);
  width: 100%;
  height: 45px;
  border-radius: 4px;
  color: white;
  font-size: 20px;
  font-weight: 600;
}

.LoginView-container .submit-button:hover {
  background-color: var(--main-color);
}
</style>
