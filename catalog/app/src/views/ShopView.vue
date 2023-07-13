<script setup lang="ts">
import { ref } from 'vue';
import axios from 'axios';
import { useCartStore } from '../stores/cart';
import { Strain } from '@/types/strain';
import { Producer } from '@/types/producer';

const cart = useCartStore();

const strains = ref<Strain[]>([]);
const producers = ref<Map<string, Producer>>(new Map());

axios.get('/api/strains')
  .then((response) => {
    strains.value = response.data
  })
  .catch((error) => {
    console.log(error)
  })

axios.get('/api/producers')
  .then((response) => {
    producers.value = response.data.reduce((map: Map<string, Producer>, producer: Producer) => {
      map.set(producer.id, producer);
      return map;
    }, new Map())
  })
  .catch((error) => {
    console.log(error)
  })

</script>

<template>
  <nav class="mx-auto w-full bg-sky-950 items-center justify-between p-3 lg:px-8 border-b border-green-700"
    aria-label="Global">
    <div class="w-full flex h-10 justify-between">
      <h2 class="text-3xl text-white">TerpScout</h2>
      <RouterLink :hidden="cart.count == 0" to="checkout" title="checkout"
        class="text-sm antialiased border-2 bg-white border-slate-800 hover:border-slate-400 py-2 px-2 rounded-lg">
        <span class="uppercase text-xs font-bold">Cart</span> <span
          class="font-mono rounded-full text-white bg-sky-950 px-2 py-1">{{ cart.count }}</span>
      </RouterLink>
    </div>
  </nav>
  <main>
    <!-- Your content -->
    <div class="bg-white m-auto">
      <div class="mx-auto max-w-2xl px-4 py-8 sm:px-6 md:max-w-5xl lg:px-8">
        <div class="py-2 mb-6">
          <h1 class="text-3xl font-bold leading-tight text-gray-700 mb-3">
            Available Strains
          </h1>
          <p class="text-xs text-gray-400">
            All the strains ready to order.
          </p>
        </div>
        <div class="mt-6 grid grid-cols-1 gap-x-6 gap-y-10 sm:grid-cols-2 lg:grid-cols-3 xl:gap-x-8">

          <!-- Strain -->
          <div v-for="strain in strains" :key="strain.id" class="group relative">
            <div class="aspect-h-1 aspect-w-1 w-full overflow-hidden rounded-md bg-gray-200 lg:aspect-none lg:h-80">
              <img :src="strain.images[0]" class="h-full w-full object-cover object-center lg:h-full lg:w-full">
            </div>
            <div class="mt-2 flex justify-between items-start">
              <div>
                <h3 class="text-lg text-gray-700">
                  {{ strain.name }}
                  <span class="text-xs text-gray-400"> {{ strain.category }}</span>
                </h3>
                <p class="mt-2 mb-2 text-xs text-gray-600">
                  <span>{{ strain.terpene_list?.join(", ") }}</span>
                </p>
                <table class="text-sm mb-2">
                  <tr>
                    <td>{{ strain.thc }}%</td>
                    <td>THC</td>
                  </tr>
                  <tr>
                    <td>{{ strain.terpenes }}%</td>
                    <td>Terpenes</td>
                  </tr>
                  <tr>
                    <td>{{ strain.total_cannabinoids }}%</td>
                    <td>Total Cannabinoids</td>
                  </tr>
                  <tr>
                    <td>{{ strain.cbd }}%</td>
                    <td>CBD</td>
                  </tr>
                </table>
              </div>
              <div>
                <h2>{{ producers.get(strain.producer_id)?.name }}</h2>
                <p class="text-sm text-right text-gray-900">${{ strain.price }}</p>
              </div>
            </div>
            <div class="mb-2 flex justify-between">
              <p class="mt-1 text-sm text-gray-900 text-ellipsis overflow-hidden whitespace-nowrap">{{ strain.genetics }}
              </p>
            </div>
            <button v-if="!cart.has(strain.id)" @click="cart.add(strain)"
              class="bg-white hover:bg-sky-950 hover:text-white border-sky-950 border rounded-md px-3 py-2 w-full">
              Add to Cart
            </button>
            <div v-else class="flex">
              <select v-model="cart.get(strain.id).quantity"
                class="mr-4 bg-white border border-sky-950 text-gray-900 text-sm rounded-lg focus:ring-sky-950 focus:border-sky-950 block w-full p-2.5 text-center">
                <option v-for="num in 10" :key="num" :value="num">{{ num }}</option>
              </select>
              <button @click="cart.del(strain.id)"
                class="text-blue-400 hover:text-blue-600 rounded-md py-2 text-sm">clear</button>
            </div>
          </div>

        </div>
      </div>
    </div>

  </main>
</template>
