import Vue from 'vue';
import Router from 'vue-router';
import Home from './../views/Home.vue';

Vue.use(Router);

export default new Router({
    routes: [
        {
            path: '/',
            name: 'home',
            component: Home
        },
        {
            path:'/authorList/:name',
            name: 'itemAuthor',
            component: ()=>import("../views/ItemAuthor")
        },
        {
            path: "/author/:id",
            name: "author",
            component: () => import('./../views/Author.vue')
        },
        {
            path:"/institution/:id",
            name:"institution",
            component: () => import('./../views/institution'),

        },
        {
            path:"/bioentity/:id",
            name:'bioentity',
            component: ()=>import('../views/Bioentity'),
        }
    ]


})

