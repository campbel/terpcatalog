<script setup lang="ts">
import { ref } from 'vue';
import axios from 'axios';
import { useCartStore } from '@/stores/cart';
import { Strain } from '@/types/strain';
import { Producer } from '@/types/producer';
import { ShoppingBagIcon } from '@heroicons/vue/24/outline'

const cart = useCartStore();

const strains = ref<Strain[]>([]);
const producers = ref<Map<string, Producer>>(new Map());
const strainsByProducer = ref<Map<string, Strain[]>>(new Map());

axios.get<Strain[]>('/api/strains')
  .then((response) => {
    strains.value = response.data;
    strainsByProducer.value = response.data.reduce((acc, strain) => {
      let strainList = acc.get(strain.producer_id)
      if (!strainList) {
        strainList = new Array<Strain>();
      }
      strainList.push(strain);
      acc.set(strain.producer_id, strainList);
      return acc;
    }, new Map<string, Strain[]>());
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
    <div class="w-full flex items-center justify-between">
      <RouterLink to="/" class="text-4xl font-bold leading-tight text-gray-100">TerpScout</RouterLink>
      <RouterLink :hidden="cart.count == 0" to="checkout"
        class="flex items-center border bg-slate-900 border-slate-950 hover:border-slate-400 px-2 py-1 rounded-md">
        <ShoppingBagIcon class="inline-block w-6 h-6 text-white" />
        <div class="w-2"></div>
        <div class="text-white">{{ cart.count }}</div>
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

        <div v-for="([id, strains]) in strainsByProducer" :key="id">
          <h2 class="text-lg font-bold uppercase">{{ producers.get(id)?.name }}</h2>
          <div class="mt-6 grid grid-cols-1 gap-x-6 gap-y-10 sm:grid-cols-2 lg:grid-cols-3 xl:gap-x-8">

            <!-- Strain -->

            <div v-for="strain in strains" :key="strain.id" class="group relative">
              <div class="aspect-h-1 aspect-w-1 w-full overflow-hidden rounded-md bg-gray-200 lg:aspect-none lg:h-80">
                <img :src="strain.images[0]" class="h-full w-full object-cover object-center lg:h-full lg:w-full">
              </div>

              <div class="mb-2">
                <h3 class="mb-1 text-md font-bold text-gray-700">
                  {{ strain.name }}
                  <span class="text-xs font-normal text-gray-400"> {{ strain.category }}</span>
                </h3>
                <p class="mb-2 text-xs text-gray-600">
                  <span v-if="strain.terpene_list && strain.terpene_list.length > 0">{{ strain.terpene_list?.join(", ")
                  }}</span>
                  <span v-else>No terpenes specified</span>
                </p>
                <p class="mb-1 text-xs text-gray-900">
                  <span class="font-bold">THC:</span> {{ strain.thc }}%
                  <span class="font-bold">CBD:</span> {{ strain.cbd }}%
                  <span class="font-bold">Terpenes:</span> {{ strain.terpenes }}%
                  <span class="font-bold" title="total cannabinoids">TC:</span> {{ strain.total_cannabinoids }}%
                </p>
                <p class="mb-1 text-sm text-gray-900">{{ strain.genetics }}</p>
                <p class="mb-1 text-sm text-gray-900">${{ strain.price }}</p>
              </div>
              <!-- <p class="mb-2 text-xs text-gray-600">
                    <span v-if="strain.terpene_list && strain.terpene_list.length > 0">{{ strain.terpene_list?.join(", ")
                    }}</span>
                    <span v-else>No terpenes specified</span>
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
                  </table> -->

              <button v-if="!cart.has(strain.id)" @click="cart.add(strain)"
                class="bg-white border-gray-200 hover:bg-sky-950 hover:text-white hover:border-sky-950 border rounded-md px-3 py-2 w-full">
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
    </div>
  </main>
</template>
