<template>
  <div class="content">
    <div class="login-container">
      <div class="login-title">Welcome Login BLog System</div>
      <div class="login-form">
        <a-form size="large" :model="loginForm" @submit="handleSubmit">
          <a-form-item field="username" hide-label hide-asterisk :rules="[{ required: true, message: 'Please input username' }]">
            <a-input v-model="loginForm.username" placeholder="Please input username">
              <template #prefix>
                <icon-user />
              </template>
            </a-input>
          </a-form-item>
          <a-form-item style="margin-top: 12px" field="password" hide-label hide-asterisk :rules="[{ required: true, message: 'Please input password' }]">
            <a-input-password v-model="loginForm.password" placeholder="Please input password" allow-clear>
              <template #prefix>
                <icon-lock />
              </template>
            </a-input-password>
          </a-form-item>
          <a-form-item field="remember_me" hide-label hide-asterisk>
            <a-checkbox style="margin-left: auto" v-model="loginForm.remember_me" value="boolean">Remember Me</a-checkbox>
          </a-form-item>
          <a-form-item hide-label hide-asterisk>
            <a-button :loading="submitLoading" type="primary" style="width: 100%" html-type="submit">Login</a-button>
          </a-form-item>
        </a-form>
      </div>
    </div>
  </div>
</template>

<script setup>

import {reactive, ref} from "vue";
import { LOGIN } from "@/api/token.js";

const loginForm = reactive({
  username: '',
  password: '',
  remember_me: false,
})

const submitLoading = ref(false)

const handleSubmit = async (data) => {
  if (data.errors === undefined) {
    try {
      submitLoading.value = true
      await LOGIN(data.values)
    } catch (e) {
      console.error(e)
    } finally {
      submitLoading.value = false
    }
  }
}

</script>

<style lang="css" scoped>

.content {
  height: 100%;
  width: 100%;
  justify-content: center;
  align-items: center;
  display: flex;
  background-color: rgb(228, 224, 224);
}

.login-container {
  width: 460px;
  height: 320px;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  background-color: rgba(255, 255, 255, 0.3);
  border-radius: 15px;
  backdrop-filter: blur(15px);
  box-shadow: 0 4px 30px rgba(0, 0, 0, 0.2);
}

.login-title {
  font-size: 20px;
  font-weight: bold;
  color: var(--color-neutral-8);
  margin: 12px 0;
}

.login-form {
  width: 100%;
  height: 100%;
  padding: 20px;
}

</style>
