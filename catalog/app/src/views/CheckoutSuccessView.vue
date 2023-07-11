<script setup lang="ts">
import { RouterLink } from 'vue-router'
</script>

<template>
    <div class="bg-white py-24 sm:py-32">
        <div class="mx-auto max-w-7xl px-6 lg:px-8">
            <div class="mx-auto max-w-2xl lg:text-center">
                <p class="mt-2 text-3xl font-bold tracking-tight text-gray-900 sm:text-4xl">Your Order has been Submitted!
                </p>
                <p class="mt-6 mb-6 text-lg leading-8 text-gray-600">Thanks for submitting your order. We'll reach out shortly
                    with an update.</p>

                <RouterLink to="/shop" class="rounded-md bg-indigo-600 px-3.5 py-2.5 text-sm font-semibold text-white shadow-sm hover:bg-indigo-500 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-indigo-600">
                    Return to the Shop
                </RouterLink>
            </div>
        </div>
    </div>
</template>
  
  