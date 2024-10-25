<template>
  <div class="container mt-5">
    <h2>My Tasks</h2>
    <form @submit.prevent="createTodo">
      <div class="mb-3">
        <label for="task" class="form-label">New Task Title</label>
        <input type="text" v-model="newTaskTitle" class="form-control" id="task" required />
      </div>
      <div class="mb-3">
        <label for="description" class="form-label">Description</label>
        <input type="text" v-model="newTaskDescription" class="form-control" id="description" required />
      </div>
      <div class="mb-3">
        <label for="deadline" class="form-label">Deadline</label>
        <input type="datetime-local" v-model="newTaskDeadline" class="form-control" id="deadline" required />
      </div>
      <button type="submit" class="btn btn-primary">Add Task</button>
    </form>

    <h3 class="mt-4">Tasks</h3>
    <ul class="list-group">
      <li
          v-for="task in todos"
          :key="task.id"
          class="list-group-item d-flex justify-content-between align-items-center"
      >
        <span>
          <strong>{{ task.title }}</strong> - {{ task.description }}
          <span v-if="task.is_done" class="badge bg-success">Done</span>
          <br />
          <small>Deadline: {{ formatDate(task.deadline) }}</small> <!-- Форматируем дату здесь -->
        </span>
        <div>
          <button class="btn btn-success btn-sm" @click="markAsDone(task)">Done</button>
          <button class="btn btn-warning btn-sm" @click="editTodo(task)">Edit</button>
          <button class="btn btn-danger btn-sm" @click="deleteTodo(task.id)">Delete</button>
        </div>
      </li>
    </ul>

    <div v-if="editMode" class="mt-4">
      <h4>Edit Task</h4>
      <form @submit.prevent="updateTodo">
        <input type="text" v-model="currentTask.title" required />
        <input type="text" v-model="currentTask.description" required />
        <input type="datetime-local" v-model="currentTask.deadline" required />
        <button type="submit" class="btn btn-primary">Update Task</button>
      </form>
    </div>
  </div>
</template>

<script>
import axios from 'axios';

export default {
  data() {
    return {
      todos: [],
      newTaskTitle: '',
      newTaskDescription: '',
      newTaskDeadline: '',
      editMode: false,
      currentTask: {},
    };
  },
  async created() {
    await this.fetchTasks();
  },
  methods: {
    async fetchTasks() {
      try {
        const response = await axios.get('http://localhost:8080/todo', { withCredentials: true });
        this.todos = response.data.todos; // Изменяем на получение задач из response.data.todos
      } catch (error) {
        console.error('Error fetching tasks:', error);
        alert('Failed to load tasks');
      }
    },
    async createTodo() {
      try {
        // Преобразование строки даты в объект Date
        const deadlineDate = new Date(this.newTaskDeadline);

        // Форматирование даты в ISO 8601
        const isoDeadline = deadlineDate.toISOString();

        // Обновлённый объект задачи с форматированной датой
        const newTodo = {
          title: this.newTaskTitle,
          description: this.newTaskDescription,
          is_done: false, // Добавляем поле is_done
          deadline: isoDeadline, // Используем отформатированную дату
        };

        const response = await axios.post('http://localhost:8080/todo', newTodo, {
          withCredentials: true
        });

        // Очистка полей ввода после успешного создания задачи
        this.newTaskTitle = '';
        this.newTaskDescription = '';
        this.newTaskDeadline = '';

        // Обновляем список задач
        await this.fetchTasks();
      } catch (error) {
        console.error('Error creating task:', error);
        alert('Failed to create task');
      }
    },
    async deleteTodo(id) {
      try {
        await axios.delete(`http://localhost:8080/todo/${id}`, { withCredentials: true });
        await this.fetchTasks(); // Обновляем список задач после удаления
      } catch (error) {
        console.error('Error deleting task:', error);
        alert('Failed to delete task');
      }
    },
    editTodo(task) {
      this.currentTask = { ...task }; // Копируем задачу для редактирования
      this.editMode = true; // Включаем режим редактирования
    },
    async updateTodo() {
      try {
        await axios.put('http://localhost:8080/todo', this.currentTask, { withCredentials: true });
        this.currentTask = {}; // Сбрасываем текущую задачу
        this.editMode = false; // Выключаем режим редактирования
        await this.fetchTasks(); // Обновляем список задач
      } catch (error) {
        console.error('Error updating task:', error);
        alert('Failed to update task');
      }
    },
    async markAsDone(task) {
      try {
        // Обновляем состояние выполнения задачи
        task.is_done = true;

        await axios.put('http://localhost:8080/todo', task, { withCredentials: true });
        await this.fetchTasks(); // Обновляем список задач после изменения
      } catch (error) {
        console.error('Error marking task as done:', error);
        alert('Failed to mark task as done');
      }
    },
    formatDate(dateString) {
      const options = { year: 'numeric', month: '2-digit', day: '2-digit', hour: '2-digit', minute: '2-digit', hour12: false };
      const date = new Date(dateString);
      return date.toLocaleString('en-GB', options).replace(',', ''); // Используем en-GB для нужного формата
    },
  },
};
</script>

<style scoped>
/* Добавьте ваши стили здесь, если необходимо */
</style>
