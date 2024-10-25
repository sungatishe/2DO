<template>
  <div class="container mt-5">
    <h2>Login</h2>
    <form @submit.prevent="login">
      <div class="mb-3">
        <label for="email" class="form-label">Email</label>
        <input type="email" v-model="email" class="form-control" id="email" required />
      </div>
      <div class="mb-3">
        <label for="password" class="form-label">Password</label>
        <input type="password" v-model="password" class="form-control" id="password" required />
      </div>
      <button type="submit" class="btn btn-primary">Login</button>
    </form>
    <p class="mt-3">Don't have an account? <router-link to="/register">Register</router-link></p>
  </div>
</template>

<script>
import axios from 'axios';

export default {
  data() {
    return {
      email: '',
      password: '',
    };
  },
  methods: {
    async login() {
      try {
        const response = await axios.post('http://localhost:8080/login', {
          email: this.email,
          password: this.password,
        }, {
          withCredentials: true, // Позволяет отправлять куки вместе с запросом
        });

        // Обновляем состояние авторизации в родительском компоненте
        this.$emit('login-success'); // Отправляем событие о успешном входе

        this.$router.push('/user'); // Переход на страницу после входа
      } catch (error) {
        console.error(error);
        alert('Login failed');
      }
    },
  },
};
</script>
