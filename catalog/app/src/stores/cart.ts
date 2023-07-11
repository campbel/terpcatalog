import { ref, computed } from 'vue'
import { defineStore } from 'pinia'

export const useCartStore = defineStore('cart', () => {
  const cart = ref<Map<string, number>>(new Map())

  const count = computed(() => {
    let count = 0;
    for (let value of cart.value.values()) {
      count += value;
    }
    return count;
  })

  function add(id: string) {
    if (cart.value.has(id)) {
      cart.value.set(id, cart.value.get(id)! + 1)
    } else {
      cart.value.set(id, 1)
    }
  }

  function set(id: string, count: number) {
    cart.value.set(id, count)
  }

  return { add, set, count }
})
