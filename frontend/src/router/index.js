import { createRouter, createWebHistory } from 'vue-router';
import Login from '../components/Login.vue';
import Register from '../components/Register.vue';
import UserProfile from "@/components/UserProfile.vue";
import Todo from "@/components/Todo.vue";
import UpdateUser from "@/components/UpdateUser.vue";

const routes = [
    { path: '/login', component: Login },
    { path: '/register', component: Register },
    {
        path: '/user',
        name: 'UserProfile',
        component: UserProfile,
    },
    { path: '/todo', component: Todo },
    { path: '/user-update', component: UpdateUser}
];

const router = createRouter({
    history: createWebHistory(),
    routes,
});

export default router;
