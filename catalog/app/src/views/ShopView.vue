<script setup lang="ts">
import { ref } from 'vue';
import axios from 'axios';
import { useCartStore } from '../stores/cart';
import { Strain } from '@/types/strain';

const cart = useCartStore();

const strains = ref<Strain[]>([]);
const producers = ref([]);

axios.get('/api/strains')
  .then((response) => {
    strains.value = response.data
  })
  .catch((error) => {
    console.log(error)
  })

axios.get('/api/strains')
  .then((response) => {
    producers.value = response.data
  })
  .catch((error) => {
    console.log(error)
  })

</script>

<template>
  <main>
    <!-- Your content -->
    <div class="bg-white m-auto">
      <div class="mx-auto max-w-2xl px-4 py-8 sm:px-6 md:max-w-7xl lg:px-8">
        <div class="flex justify-between items-center py-2 px-4 h-14 shadow border rounded-md ">
          <h2 class="text-2xl font-bold tracking-tight text-gray-900">Available Strains</h2>
          <RouterLink v-if="cart.count > 0" to="checkout" title="checkout"
            class="text-sm antialiased border-2 bg-white border-slate-400 hover:border-slate-600 hover:text-white py-2 px-2 rounded-lg">
            🛒 <span class="font-mono rounded-full text-white bg-orange-500 px-2 py-1">{{ cart.count }}</span>
          </RouterLink>
        </div>
        <div class="mt-6 grid grid-cols-1 gap-x-6 gap-y-10 sm:grid-cols-2 lg:grid-cols-4 xl:gap-x-8">

          <!-- Strain -->
          <div v-for="strain in strains" :key="strain.id" class="group relative">
            <div
              class="aspect-h-1 aspect-w-1 w-full overflow-hidden rounded-md bg-gray-200 lg:aspect-none lg:h-80">
              <img :src="strain.images[0]" class="h-full w-full object-cover object-center lg:h-full lg:w-full">
            </div>
            <div class="mt-2 flex justify-between items-start">
              <div>
                <h3 class="text-lg text-gray-700">
                  {{ strain.name }}
                  <span class="text-xs text-gray-400"> {{ strain.category }}</span>
                </h3>
                <p class="mt-1 text-sm text-gray-900">
                  <span title="thc%">{{ strain.thc }}%</span> | <span title="terpenes%">{{ strain.terpenes }}%</span>
                </p>
              </div>
              <div>
                <p class="text-sm text-right font-medium text-gray-900">${{ strain.price }}</p>
              </div>
            </div>
            <div class="mb-2 flex justify-between">
              <p class="mt-1 text-sm text-gray-500 text-ellipsis overflow-hidden whitespace-nowrap">{{ strain.genetics }}</p>
            </div>
            <button v-if="!cart.has(strain.id)" @click="cart.add(strain)" class="bg-gray-200 hover:bg-gray-400 rounded-md px-3 py-2 w-full">
              Add to Cart
            </button>
            <div v-else class="flex">
              <select v-model="cart.get(strain.id).quantity" 
                class="mr-4 bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 text-center">
                <option v-for="num in 10" :key="num" :value="num">{{ num }}</option>
              </select>
              <button @click="cart.del(strain.id)" 
                class="bg-gray-200 hover:bg-gray-400 rounded-md px-3 py-2">Remove</button>
            </div>
          </div>

        </div>
      </div>
    </div>

  </main>
</template>
