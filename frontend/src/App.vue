<template>
  <div>
    <nav class="navbar navbar-expand-lg navbar-light bg-light">
      <div class="container-fluid">
        <a class="navbar-brand" href="#">2DO!</a>
        <button
            class="navbar-toggler"
            type="button"
            data-bs-toggle="collapse"
            data-bs-target="#navbarNav"
            aria-controls="navbarNav"
            aria-expanded="false"
            aria-label="Toggle navigation"
        >
          <span class="navbar-toggler-icon"></span>
        </button>
        <div class="collapse navbar-collapse" id="navbarNav">
          <ul class="navbar-nav">
            <li class="nav-item" v-if="!isAuthenticated">
              <router-link class="nav-link" to="/login">Login</router-link>
            </li>
            <li class="nav-item" v-if="!isAuthenticated">
              <router-link class="nav-link" to="/register">Register</router-link>
            </li>
            <li class="nav-item" v-if="isAuthenticated">
              <router-link class="nav-link" to="/todo">Tasks</router-link>
            </li>
            <li class="nav-item" v-if="isAuthenticated">
              <router-link class="nav-link" to="/user">User</router-link>
            </li>
            <li class="nav-item" v-if="isAuthenticated">
              <a class="nav-link" @click="logout">Logout</a>
            </li>
          </ul>
        </div>
      </div>
    </nav>
    <router-view @login-success="checkAuthorization"></router-view> <!-- Обработка события -->
  </div>
</template>

<script>
import axios from 'axios';

export default {
  data() {
    return {
      isAuthenticated: false,
    };
  },
  methods: {
    async checkAuthorization() {
      try {
        await axios.get('http://localhost:8080/user', { withCredentials: true });
        this.isAuthenticated = true; // Пользователь авторизован
      } catch (error) {
        this.isAuthenticated = false; // Ошибка авторизации
        console.error('User is not authenticated:', error);
      }
    },
    async logout() {
      try {
        await axios.post('http://localhost:8080/logout', {}, { withCredentials: true });
        this.isAuthenticated = false; // Обновляем состояние после логаута
        this.$router.push('/login'); // Перенаправляем на страницу входа
      } catch (error) {
        console.error('Logout failed:', error);
      }
    },
  },
  created() {
    this.checkAuthorization(); // Проверяем авторизацию при создании компонента
  },
};
</script>

<style scoped>
.nav-link {
  cursor: pointer;
}
</style>
