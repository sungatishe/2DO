import { createApp } from 'vue';
import App from './App.vue';
import router from './router'; // Импортируйте ваш маршрутизатор
import 'bootstrap/dist/css/bootstrap.min.css'; // Импортируйте стили Bootstrap

const app = createApp(App);


app.use(router); // Используйте маршрутизатор
app.mount('#app');
