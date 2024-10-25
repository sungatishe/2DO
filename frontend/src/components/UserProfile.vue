<template>
  <div class="container mt-5">
    <h2>User Profile</h2>
    <div v-if="user" class="card mb-3">
      <div class="row g-0">
        <div class="col-md-4">
          <img :src="user.user.avatar" alt="User Avatar" class="img-fluid rounded-start" />
        </div>
        <div class="col-md-8">
          <div class="card-body">
            <h5 class="card-title">{{ user.user.username }}</h5>
            <p class="card-text"><strong>Email:</strong> {{ user.user.email }}</p>
            <p class="card-text"><strong>Description:</strong> {{ user.user.description }}</p>
            <router-link to="/user-update">Update details</router-link>
          </div>
        </div>
      </div>
    </div>
    <div v-else>
      <p>Loading user profile...</p>
    </div>
  </div>
</template>

<script>
import axios from 'axios';

export default {
  data() {
    return {
      user: null, // Изначально user будет null
    };
  },
  async created() {
    await this.fetchUserProfile();
  },
  methods: {
    async fetchUserProfile() {
      try {
        const response = await axios.get('http://localhost:8080/user', {
          withCredentials: true, // Включаем куки в запрос
        });

        console.log('User Profile Data:', response.data); // Выводим данные в консоль
        this.user = response.data; // Сохраняем ответ в data
      } catch (error) {
        console.error('Error fetching user profile:', error);
        alert('Failed to load user profile');
      }
    },
  },
};
</script>

<style scoped>
.card {
  border: 1px solid #e0e0e0;
  border-radius: 10px;
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
}
.card-title {
  font-weight: bold;
  font-size: 1.5rem;
}
</style>
