<template>
  <div class="content">
    <div class="login-container">
      <div class="login-title">Welcome Login BLog System</div>
      <div class="login-form">
        <a-form size="large" :model="loginForm" @submit="handleSubmit">
          <a-form-item field="username" hide-label hide-asterisk>
            <a-input v-model="loginForm.username" placeholder="Please input username">
              <template #prefix>
                <icon-user />
              </template>
            </a-input>
          </a-form-item>
          <a-form-item field="password" hide-label hide-asterisk>
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
            <a-button type="primary" style="width: 100%" html-type="submit">Login</a-button>
          </a-form-item>
        </a-form>
      </div>
    </div>
  </div>
</template>

<script setup>

import {reactive} from "vue";

const loginForm = reactive({
  username: '',
  password: '',
  remember_me: false,
})

const handleSubmit = (data) => {
  console.log(data)
}

</script>

<style lang="css" scoped>

.content {
  height: 100%;
  width: 100%;
  justify-content: center;
  align-items: center;
  display: flex;
}

.login-container {
  width: 460px;
  height: 400px;
  display: flex;
  flex-direction: column;
  align-items: center;
}

.login-title {
  font-size: 20px;
  font-weight: bold;
  color: var(--color-neutral-8);
  margin-top: 12px;
}

.login-form {
  width: 100%;
  height: 100%;
  padding: 12px;
}

</style>
