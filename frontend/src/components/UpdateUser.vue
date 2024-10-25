<template>
  <div class="container mt-5">
    <h2>Update User Information</h2>
    <form v-if="user" @submit.prevent="updateUser">
      <div class="mb-3">
        <label for="username" class="form-label">Username</label>
        <input
            type="text"
            v-model="user.user.username"
        class="form-control"
        id="username"
        />
      </div>
      <div class="mb-3">
        <label for="email" class="form-label">Email</label>
        <input
            type="email"
            v-model="user.user.email"
        class="form-control"
        id="email"
        />
      </div>
      <div class="mb-3">
        <label for="description" class="form-label">Description</label>
        <textarea
            v-model="user.user.description"
        class="form-control"
        id="description"
        ></textarea>
      </div>
      <div class="mb-3">
        <label for="avatar" class="form-label">Avatar URL</label>
        <input
            type="url"
            v-model="user.user.avatar"
        class="form-control"
        id="avatar"
        />
      </div>
      <button type="submit" class="btn btn-primary">Update User</button>
    </form>
    <div v-else>
      <p>Loading user data...</p> <!-- Показать индикатор загрузки -->
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
    await this.fetchUserData(); // Загружаем данные пользователя при создании компонента
  },
  methods: {
    async fetchUserData() {
      try {
        const response = await axios.get('http://localhost:8080/user', {
          withCredentials: true, // Включаем куки в запрос
        });

        console.log('Full response from GET /user:', response.data); // Выводим полный ответ в консоль

        // Убедитесь, что данные корректно присваиваются
        if (response.data && response.data.user) {
          this.user = response.data; // Сохраняем ответ в data
          console.log('User Profile Data:', this.user); // Выводим данные пользователя в консоль
        } else {
          console.error('User data not found in response');
        }
      } catch (error) {
        console.error('Error fetching user profile:', error);
        alert('Failed to load user profile');
      }
    },
    async updateUser() {
      try {
        var userData = {
          user_id: this.user.user.id, // Используем user.id для обновления
          username: this.user.user.username,
          email: this.user.user.email,
          description: this.user.user.description,
          avatar: this.user.user.avatar,
        }
        await axios.put('http://localhost:8080/user', userData, { withCredentials: true });
        alert('User updated successfully!');
        this.$router.push('/user');
      } catch (error) {
        console.error('Error updating user:', error);
        alert('Failed to update user');
      }
    },
  },
};
</script>

<style scoped>
/* Добавьте ваши стили здесь, если необходимо */
</style>
