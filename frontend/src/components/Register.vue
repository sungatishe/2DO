<template>
  <div class="container mt-5">
    <h2>Register</h2>
    <form @submit.prevent="register">
      <div class="mb-3">
        <label for="username" class="form-label">Username</label>
        <input type="text" v-model="username" class="form-control" id="username" required />
      </div>
      <div class="mb-3">
        <label for="email" class="form-label">Email</label>
        <input type="email" v-model="email" class="form-control" id="email" required />
      </div>
      <div class="mb-3">
        <label for="password" class="form-label">Password</label>
        <input type="password" v-model="password" class="form-control" id="password" required />
      </div>
      <button type="submit" class="btn btn-primary">Register</button>
    </form>
    <p class="mt-3">Already have an account? <router-link to="/login">Login</router-link></p>
  </div>
</template>

<script>
import axios from 'axios';

export default {
  data() {
    return {
      username: '', // Добавляем поле для имени пользователя
      email: '',
      password: '',
    };
  },
  methods: {
    async register() {
      try {
        await axios.post('http://localhost:8080/register', {
          username: this.username, // Передаем имя пользователя
          email: this.email,
          password: this.password,
        });
        this.$router.push('/login'); // Переход на страницу входа после успешной регистрации
      } catch (error) {
        console.error(error);
        alert('Registration failed');
      }
    },
  },
};
</script>

<style scoped>
/* Добавьте ваши стили здесь, если необходимо */
</style>
