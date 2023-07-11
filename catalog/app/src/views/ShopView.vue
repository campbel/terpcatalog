<script setup lang="ts">
import { ref } from 'vue';
import axios from 'axios';

class Strain {
  id: string = '';
  name: string = '';
  category: string = '';
  genetics: string = '';
  thc: number = 0;
  terpenes: number = 0;
  price: number = 0;
  images: string[] = [];
}

const strains = ref<Strain[]>([]);
const producers = ref([]);
const cart = ref<Map<string,number>>(new Map());

axios.get('/api/strains')
  .then((response) => {
    console.log(response.data)
    strains.value = response.data
  })
  .catch((error) => {
    console.log(error)
  })

axios.get('/api/strains')
  .then((response) => {
    console.log(response.data)
    producers.value = response.data
  })
  .catch((error) => {
    console.log(error)
  })

function addToCart(strain: Strain) {
  if (cart.value.has(strain.id))
    cart.value.set(strain.id, cart.value.get(strain.id)! + 1);
  else
    cart.value.set(strain.id, 1);
}

function itemCount() {
  let count = 0;
  for (let value of cart.value.values()) {
    count += value;
  }
  return count;
}

</script>

<template>
  <main>
    <div class="mx-auto max-w-7xl py-6 sm:px-6 lg:px-8">
        <!-- Your content -->
      <div class="bg-white m-auto">
        <div class="mx-auto max-w-2xl px-4 py-16 sm:px-6 sm:py-24 lg:max-w-7xl lg:px-8">
          <div class="flex justify-between items-center py-2 px-4 h-14 shadow border rounded-md ">
            <h2 class="text-2xl font-bold tracking-tight text-gray-900">Available Strains</h2>
            <button v-if="itemCount() > 0" title="checkout"
                class="text-sm antialiased border-2 bg-white border-slate-400 hover:border-slate-600 hover:text-white py-2 px-2 rounded-lg">
              🛒 <span class="font-mono rounded-full text-white bg-orange-500 px-2 py-1">{{ itemCount() }}</span>
            </button>
          </div>
          <div class="mt-6 grid grid-cols-1 gap-x-6 gap-y-10 sm:grid-cols-2 lg:grid-cols-4 xl:gap-x-8">

            <!-- Strain -->
            <div v-for="strain in strains" :key="strain.id"
              class="group relative">
              <div class="aspect-h-1 aspect-w-1 w-full overflow-hidden rounded-md bg-gray-200 lg:aspect-none group-hover:opacity-75 lg:h-80">
                <img :src="strain.images[0]" class="h-full w-full object-cover object-center lg:h-full lg:w-full">
              </div>
              <div class="mt-2 flex justify-between items-end">
                <div>
                  <h3 class="text-lg text-gray-700">
                    {{ strain.name }}
                    <span class="text-xs text-gray-400"> {{ strain.category }}</span>
                  </h3>
                </div>
                <div>
                  <p class="text-sm text-right font-medium text-gray-900">${{ strain.price }}</p>
                </div>
              </div>
              <div class="mb-2 flex justify-between">
                <p class="mt-1 text-sm text-gray-500">{{ strain.genetics }}</p>
                <p class="mt-1 text-sm text-gray-900">
                  <span title="thc%">{{ strain.thc }}%</span> | <span title="terpenes%">{{ strain.terpenes }}%</span></p>
              </div>
              <button @click="addToCart(strain)"
                class="bg-gray-200 hover:bg-gray-400 rounded-md px-3 py-2 w-full">
                Add to Cart
              </button>
            </div>

          </div>
        </div>
      </div>

    </div>
  </main>
</template>
