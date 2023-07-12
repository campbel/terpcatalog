<script setup lang="ts">
import { ref } from 'vue';
import axios from 'axios';
import { useCartStore } from '../stores/cart';
import { useRouter } from 'vue-router';

const cart = useCartStore();
const router = useRouter();

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

class Order {
  strain: Strain = new Strain();
  quantity: number = 0;
}

const strains = ref<Map<string, Strain>>(new Map());
const orders = ref<Order[]>([]);

axios.get('/api/strains')
  .then((response) => {
    strains.value = response.data.reduce((map: Map<string, Strain>, strain: Strain) => {
      map.set(strain.id, strain);
      return map;
    }, new Map())
    orders.value = cart.items.map((item) => {
      return {
        strain: strains.value.get(item.id) as Strain,
        quantity: item.count
      }
    })
  })
  .catch((error) => {
    console.log(error)
  })

function onSubmit() {
  cart.reset();
  router.push({ name: 'checkout-success' });
}

</script>

<template>
  <main>
    <!-- Your content -->
    <div class="bg-white m-auto">
      <div class="mx-auto max-w-2xl px-4 py-8 sm:px-6 md:max-w-5xl lg:px-8">
        <div class="flex justify-between items-center py-2 px-4 h-14 shadow border rounded-md ">
          <h2 class="text-2xl font-bold tracking-tight text-gray-900">Checkout</h2>

        </div>

        <div class="mt-6 md:flex justify-between">

          <div class="md:w-1/2 w-full pr-6 mb-6">
            <form @submit.prevent="onSubmit">
              <div class="space-y-12">

                <div class="border-b border-gray-900/10 pb-12">
                  <h2 class="text-base font-semibold leading-7 text-gray-900">Personal Information</h2>
                  <p class="mt-1 text-sm leading-6 text-gray-600">Use a permanent address where you can receive mail.</p>

                  <div class="mt-10 grid grid-cols-1 gap-x-6 gap-y-8 md:grid-cols-6">

                    <div class="md:col-span-3">
                      <label for="first-name" class="block text-sm font-medium leading-6 text-gray-900">First name</label>
                      <div class="mt-2">
                        <input type="text" name="first-name" id="first-name" autocomplete="given-name"
                          class="block w-full rounded-md border-0 p-1.5 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6" />
                      </div>
                    </div>

                    <div class="md:col-span-3">
                      <label for="last-name" class="block text-sm font-medium leading-6 text-gray-900">Last name</label>
                      <div class="mt-2">
                        <input type="text" name="last-name" id="last-name" autocomplete="family-name"
                          class="block w-full rounded-md border-0 p-1.5 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6" />
                      </div>
                    </div>

                    <div class="md:col-span-3">
                      <label for="email" class="block text-sm font-medium leading-6 text-gray-900">Email address</label>
                      <div class="mt-2">
                        <input id="email" name="email" type="email" autocomplete="email"
                          class="block w-full rounded-md border-0 p-1.5 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6" />
                      </div>
                    </div>

                    <div class="md:col-span-3">
                      <label for="phone" class="block text-sm font-medium leading-6 text-gray-900">Phone</label>
                      <div class="mt-2">
                        <input id="phone" name="phone" type="phone" autocomplete="phone"
                          class="block w-full rounded-md border-0 p-1.5 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6" />
                      </div>
                    </div>

                    <div class="col-span-full">
                      <label for="street-address" class="block text-sm font-medium leading-6 text-gray-900">Street
                        address</label>
                      <div class="mt-2">
                        <input type="text" name="street-address" id="street-address" autocomplete="street-address"
                          class="block w-full rounded-md border-0 p-1.5 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6" />
                      </div>
                    </div>

                    <div class="md:col-span-2">
                      <label for="city" class="block text-sm font-medium leading-6 text-gray-900">City</label>
                      <div class="mt-2">
                        <input type="text" name="city" id="city" autocomplete="address-level2"
                          class="block w-full rounded-md border-0 p-1.5 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6" />
                      </div>
                    </div>

                    <div class="md:col-span-2">
                      <label for="region" class="block text-sm font-medium leading-6 text-gray-900">State</label>
                      <div class="mt-2">
                        <input type="text" name="region" id="region" autocomplete="address-level1"
                          class="block w-full rounded-md border-0 p-1.5 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6" />
                      </div>
                    </div>

                    <div class="md:col-span-2">
                      <label for="postal-code" class="block text-sm font-medium leading-6 text-gray-900">ZIP Code</label>
                      <div class="mt-2">
                        <input type="text" name="postal-code" id="postal-code" autocomplete="postal-code"
                          class="block w-full rounded-md border-0 p-1.5 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6" />
                      </div>
                    </div>
                  </div>
                </div>


              </div>

              <div class="mt-6 flex items-center justify-end gap-x-6">
                <RouterLink to="shop" type="button" class="text-sm font-semibold leading-6 text-gray-900">Cancel
                </RouterLink>
                <button type="submit"
                  class="rounded-md bg-indigo-600 px-3 py-2 text-sm font-semibold text-white shadow-sm hover:bg-indigo-500 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-indigo-600">Place
                  Order</button>
              </div>
            </form>
          </div>

          <div class="md:w-1/2 w-full pl-6 mb-6">
            <h2 class="text-base font-semibold leading-7 text-gray-900">Order Summary</h2>

            <div v-for="(order, index) in orders" class="border-b py-4">
              <div class="flex mb-3 items-end">
                <img :src="order.strain.images[0]" class="mr-3 h-32 rounded-lg" />
                <div>
                  <h2>{{ order.strain.name }}<p class="text-xs">{{ order.strain.category }}</p>
                  </h2>
                  <p>${{ order.strain.price }} x {{ order.quantity }} = ${{ order.quantity * order.strain.price }}</p>
                </div>
              </div>
              <div class="flex">
                <select v-model="order.quantity"
                  class="mr-4 bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 text-center">
                  <option v-for="num in 10" :key="num" :value="num">{{ num }}</option>
                </select>
                <button @click="orders.splice(index, 1)" class="bg-gray-200 hover:bg-gray-400 rounded-md px-3 py-2">
                  Remove</button>
              </div>
            </div>

            <div class="flex justify-between mt-3">
              <h2 class="bold text-xl">Total</h2>
              <h2 class="bold text-xl">${{ orders.reduce((acc, order) => acc + (order.quantity * order.strain.price), 0) }}</h2>
            </div>

          </div>

        </div>
      </div>
    </div>

  </main>
</template>