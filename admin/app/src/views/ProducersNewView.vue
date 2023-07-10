<script setup lang="ts">
import { ref } from 'vue'
import axios from 'axios';
import { useRouter } from 'vue-router';

const router = useRouter();

const images: [] = []
const newProducer = ref({
  "name": "",
  "email": "",
  "website": "",
  "address": "",
  "contact": {
    "firstName": "",
    "lastName": "",
    "email": "",
    "phone": "",
  },
  "phone": "",
})

function onSubmit() {
  axios.post("/api/producers", newProducer.value)
    .then((response) => {
      router.push({ name: "producers" });
    })
    .catch((error) => {
      console.log(error);
    });
}

</script>

<template>
  <main>
    <div class="w-full">
      <div class="bg-white max-w-lg mx-auto shadow-md rounded-xl border border-gray-200or px-8 pt-6 pb-8 mb-4">

        <h1 class="text-3xl space-y-1 font-bold leading-tight text-gray-700 mb-3">
          Add a Producer
        </h1>

        <hr class="h-px my-6 bg-gray-200 border-0 dark:bg-gray-200">

        <form @submit.prevent="onSubmit" class="">

          <!-- Row 1-->
          <div class="flex flex-wrap -mx-3 mb-3">

            <!-- Name -->
            <div class="w-full px-3 mb-3 md:mb-0">
              <label class="block uppercase tracking-wide text-gray-700 text-xs font-bold mb-2" for="grid-name">
                Name
              </label>
              <input v-model="newProducer.name" id="grid-name" type="text" placeholder="Terp Farms" required
                class="appearance-none block w-full bg-gray-200 text-gray-700 border border-gray-200 rounded py-3 px-4 mb-3 leading-tight focus:outline-none focus:bg-white">
            </div>

            <!-- Website -->
            <div class="w-full px-3 mb-3 md:mb-0">
              <label class="block uppercase tracking-wide text-gray-700 text-xs font-bold mb-2" for="grid-website">
                Website
              </label>
              <input v-model="newProducer.website" id="grid-website" type="text" placeholder="terpfarms.com" required
                class="appearance-none block w-full bg-gray-200 text-gray-700 border border-gray-200 rounded py-3 px-4 mb-3 leading-tight focus:outline-none focus:bg-white">
            </div>

            <!-- Email -->
            <div class="w-full md:w-1/2 px-3 mb-3 md:mb-0">
              <label class="block uppercase tracking-wide text-gray-700 text-xs font-bold mb-2" for="grid-email">
                Email
              </label>
              <input v-model="newProducer.email" id="grid-email" type="text" placeholder="contact@terpfarms.com" required
                class="appearance-none block w-full bg-gray-200 text-gray-700 border border-gray-200 rounded py-3 px-4 mb-3 leading-tight focus:outline-none focus:bg-white">
            </div>

            <!-- Phone -->
            <div class="w-full md:w-1/2 px-3 mb-3 md:mb-0">
              <label class="block uppercase tracking-wide text-gray-700 text-xs font-bold mb-2" for="grid-phone">
                Phone
              </label>
              <input v-model="newProducer.phone" id="grid-phone" type="text" placeholder="555-555-5555" required
                class="appearance-none block w-full bg-gray-200 text-gray-700 border border-gray-200 rounded py-3 px-4 mb-3 leading-tight focus:outline-none focus:bg-white">
            </div>

            <!-- Address -->
            <div class="w-full px-3 mb-3 md:mb-0">
              <label class="block uppercase tracking-wide text-gray-700 text-xs font-bold mb-2" for="grid-address">
                Address
              </label>
              <textarea v-model="newProducer.address" id="grid-address" type="text" placeholder="123 Farm Rd. Portland" required
                class="appearance-none block w-full bg-gray-200 text-gray-700 border border-gray-200 rounded py-3 px-4 mb-3 leading-tight focus:outline-none focus:bg-white">
                </textarea>
            </div>

            </div>

          <hr class="h-px my-6 bg-gray-200 border-0 dark:bg-gray-200">

          <h2 class="text-xl space-y-1 font-bold leading-tight text-gray-700 mb-3">
          Contact
          </h2>
          <p class="text-xs text-gray-400 mb-3">
            Information about the primary contact for the producer.
          </p>

          <div class="flex flex-wrap -mx-3 mb-3">

            <!-- Contact First Name -->
            <div class="w-full md:w-1/2 px-3 mb-3 md:mb-0">
              <label class="block uppercase tracking-wide text-gray-700 text-xs font-bold mb-2" for="grid-contact-first">
                First Name
              </label>
              <input v-model="newProducer.contact.firstName" id="grid-phone" type="text" placeholder="Sarah" required
                class="appearance-none block w-full bg-gray-200 text-gray-700 border border-gray-200 rounded py-3 px-4 mb-3 leading-tight focus:outline-none focus:bg-white">
            </div>

            <!-- Contact Last Name -->
            <div class="w-full md:w-1/2 px-3 mb-3 md:mb-0">
              <label class="block uppercase tracking-wide text-gray-700 text-xs font-bold mb-2" for="grid-contact-last">
                Last Name
              </label>
              <input v-model="newProducer.contact.lastName" id="grid-contact-last" type="text" placeholder="Smith" required
                class="appearance-none block w-full bg-gray-200 text-gray-700 border border-gray-200 rounded py-3 px-4 mb-3 leading-tight focus:outline-none focus:bg-white">
            </div>

            <!-- Contact Email -->
            <div class="w-full md:w-1/2 px-3 mb-3 md:mb-0">
              <label class="block uppercase tracking-wide text-gray-700 text-xs font-bold mb-2" for="grid-contact-email">
                Email
              </label>
              <input v-model="newProducer.contact.email" id="grid-contact-email" type="text" placeholder="sarah@terpfarms.com" required
                class="appearance-none block w-full bg-gray-200 text-gray-700 border border-gray-200 rounded py-3 px-4 mb-3 leading-tight focus:outline-none focus:bg-white">
            </div>

            <!-- Contact Phone -->
            <div class="w-full md:w-1/2 px-3 mb-3 md:mb-0">
              <label class="block uppercase tracking-wide text-gray-700 text-xs font-bold mb-2" for="grid-contact-first">
                Phone
              </label>
              <input v-model="newProducer.contact.phone" id="grid-phone" type="text" placeholder="555-555-5555" required
                class="appearance-none block w-full bg-gray-200 text-gray-700 border border-gray-200 rounded py-3 px-4 mb-3 leading-tight focus:outline-none focus:bg-white">
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