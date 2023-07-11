import { ref, computed } from 'vue'
import { defineStore } from 'pinia'

export const useCartStore = defineStore('cart', () => {
  const cart = ref<Map<string, number>>(new Map())

  const items = computed(() => {
    const items: { id: string; count: number }[] = []
    for (let [id, count] of cart.value.entries()) {
      items.push({ id, count })
    }
    return items
  })

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

  function has(id: string) {
    return cart.value.has(id)
  }

  function set(id: string, count: number) {
    cart.value.set(id, count)
  }

  function del(id: string) {
    if (cart.value.has(id)) {
      cart.value.delete(id)
    }
  }

  return { add, has, set, del, items, count, cart }
})
