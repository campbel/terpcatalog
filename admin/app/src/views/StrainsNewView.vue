<script setup lang="ts">
import { ref } from 'vue'
import axios from 'axios';
import { useRouter } from 'vue-router';

const router = useRouter();

const categories = ref([
  "Indica",
  "Sativa",
  "Hybrid",
])

const images: [] = []
const newStrain = ref({
  "name": "",
  "category": "",
  "genetics": "",
  "thc": 0,
  "terpenes": 0,
  "price": 0,
  "harvest_date": "",
  "images": new Array<string>(),
})

function onSubmit() {
  axios.post("/api/strains", newStrain.value)
    .then((response) => {
      router.push({ name: "strains" });
    })
    .catch((error) => {
      console.log(error);
    });
}

function uploadImage(e: any) {
  if (!e.target || !e.target.files || e.target.files.length == 0) return
  const image = e.target.files[0];
  const reader = new FileReader();
  reader.readAsDataURL(image);
  reader.onload = e => {
    if (!e.target) return
    let image = e.target.result;
    if (!image) return;
    if (typeof image === 'string') newStrain.value.images.push(image);
  };
}

function removeImage(index: number) {
  newStrain.value.images.splice(index, 1);
}
</script>

<template>
  <main>
    <div class="w-full">
      <div class="bg-white max-w-lg mx-auto shadow-md rounded px-8 pt-6 pb-8 mb-4">

        <h1 class="text-3xl space-y-1 font-bold leading-tight text-gray-700 mb-3">
          Add a Strain
        </h1>

        <hr class="h-px my-6 bg-gray-200 border-0 dark:bg-gray-200">

        <form @submit.prevent="onSubmit" class="">

          <!-- Row 1-->
          <div class="flex flex-wrap -mx-3 mb-3">

            <!-- Name -->
            <div class="w-full md:w-1/2 px-3 mb-3 md:mb-0">
              <label class="block uppercase tracking-wide text-gray-700 text-xs font-bold mb-2" for="grid-name">
                Name
              </label>
              <input v-model="newStrain.name" id="grid-name" type="text" placeholder="Green Blaze" required
                class="appearance-none block w-full bg-gray-200 text-gray-700 border border-gray-200 rounded py-3 px-4 mb-3 leading-tight focus:outline-none focus:bg-white">
            </div>

            <!-- Category -->
            <div class="w-full md:w-1/2 px-3 mb-3 md:mb-0">
              <label class="block uppercase tracking-wide text-gray-700 text-xs font-bold mb-2" for="grid-category">
                Category
              </label>
              <div class="relative">
                <select v-model="newStrain.category" id="grid-category" required
                  class="block appearance-none w-full bg-gray-200 border border-gray-200 text-gray-700 py-3 px-4 pr-8 rounded leading-tight focus:outline-none focus:bg-white focus:border-gray-500">
                  <option v-for="option in categories" :key="option" :value="option">
                    {{ option }}
                  </option>
                </select>
                <div class="pointer-events-none absolute inset-y-0 right-0 flex items-center px-2 text-gray-700">
                  <svg class="fill-current h-4 w-4" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 20 20">
                    <path d="M9.293 12.95l.707.707L15.657 8l-1.414-1.414L10 10.828 5.757 6.586 4.343 8z" />
                  </svg>
                </div>
              </div>
            </div>

          </div>

          <!-- Row 2-->
          <div class="flex flex-wrap -mx-3 mb-3">

            <div class="w-full md:w px-3 mb-3 md:mb-0">
              <label class="block uppercase tracking-wide text-gray-700 text-xs font-bold mb-2" for="grid-genetics">
                Genetics
              </label>
              <textarea v-model="newStrain.genetics" id="grid-genetics" required
                class="appearance-none block w-full bg-gray-200 text-gray-700 border border-gray-200 rounded py-3 px-4 mb-3 leading-tight focus:outline-none focus:bg-white focus:border-gray-500"></textarea>
            </div>

          </div>

          <!-- Row 3-->
          <div class="flex flex-wrap -mx-3 mb-3">

            <!-- THC -->
            <div class="w-full md:w-1/2 px-3 mb-3 md:mb-0">
              <label class="block uppercase tracking-wide text-gray-700 text-xs font-bold mb-2" for="grid-thc">
                THC% (0-100)
              </label>
              <input v-model.number="newStrain.thc" id="grid-thc" type="number" placeholder="25.0" required min="0"
                max="100"
                class="appearance-none block w-full bg-gray-200 text-gray-700 border border-gray-200 rounded py-3 px-4 mb-3 leading-tight focus:outline-none focus:bg-white focus:border-gray-500">
            </div>

            <!-- Terpenes -->
            <div class="w-full md:w-1/2 px-3 mb-3 md:mb-0">
              <label class="block uppercase tracking-wide text-gray-700 text-xs font-bold mb-2" for="grid-terpenes">
                Terpenes (0-100)
              </label>
              <input v-model.number="newStrain.terpenes" id="grid-terpenes" type="number" placeholder="25.0" required
                min="0" max="100"
                class="appearance-none block w-full bg-gray-200 text-gray-700 border border-gray-200 rounded py-3 px-4 mb-3 leading-tight focus:outline-none focus:bg-white focus:border-gray-500">
            </div>

          </div>

          <!-- Row 4-->
          <div class="flex flex-wrap -mx-3 mb-3">

            <!-- Price -->
            <div class="w-full md:w-1/2 px-3 mb-3 md:mb-0">
              <label class="block uppercase tracking-wide text-gray-700 text-xs font-bold mb-2" for="grid-price">
                Price (USD)
              </label>
              <input v-model.number="newStrain.price" id="grid-price" type="number" placeholder="25.0" required min="0"
                class="appearance-none block w-full bg-gray-200 text-gray-700 border border-gray-200 rounded py-3 px-4 mb-3 leading-tight focus:outline-none focus:bg-white focus:border-gray-500">
            </div>

            <!-- Harvest Date -->
            <div class="w-full md:w-1/2 px-3 mb-3 md:mb-0">
              <label class="block uppercase tracking-wide text-gray-700 text-xs font-bold mb-2" for="grid-harvest-date">
                Harvest Date
              </label>
              <input v-model="newStrain.harvest_date" id="grid-harvest-date" type="date" placeholder="25.0" required
                class="appearance-none block w-full bg-gray-200 text-gray-700 border border-gray-200 rounded py-3 px-4 mb-3 leading-tight focus:outline-none focus:bg-white focus:border-gray-500">
            </div>

          </div>

          <hr class="h-px my-6 bg-gray-200 border-0 dark:bg-gray-200">

          <!-- Row 5 -->


          <div class="flex flex-wrap -mx-3 mb-3 justify-between items-center">

            <div class="px-3 mb-3">
              <h2 class="block text-xl space-y-1 font-bold leading-tight text-gray-700 mb-3">Images</h2>
              <p class="text-sm text-gray-400 mb-3">Upload images of the product here.</p>
            </div>

            <!-- Image -->
            <div class="px-3 mb-3">
              <label class="block  cursor-pointer bg-slate-400 hover:bg-slate-700 text-white font-bold py-2 px-4 rounded focus:outline-none focus:shadow-outline" for="grid-photo">
                Add Image
                <input type="file" accept="image/jpeg;image/png" id="grid-photo" @change="uploadImage" />
              </label>
            </div>

          </div>

          <div class="flex flex-wrap -mx-3 mb-3">

            <div v-for="(image, index) in newStrain.images" class="md:w-1/2 px-3 mb-3">
              <img :src="image" class="rounded shadow-md border-gray-200 max-h-48 border"/>
              <button type="button" @click="removeImage(index)" class="uppercase mt-2 text-xs text-gray-300 hover:text-gray-500">
                Remove
              </button>
            </div>

          </div>


          <hr class="h-px my-6 bg-gray-200 border-0 dark:bg-gray-200">

          <!-- Footer -->
          <div class="flex flex-wrap justify-end -mx-3 mb-0">

            <!-- Submit -->
            <div class="px-3 mb-3 md:mb-0">
              <button
                class="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded focus:outline-none focus:shadow-outline"
                type="submit">
                Submit
              </button>
            </div>

          </div>

        </form>
      </div>
    </div>
  </main>
</template>

<style>
input[type="file"] {
  display: none;
}
</style>