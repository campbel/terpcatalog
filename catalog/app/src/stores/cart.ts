import { ref, computed } from 'vue'
import { defineStore } from 'pinia'
import { OrderItem } from '@/types/order'
import { Strain } from '@/types/strain'

export const useCartStore = defineStore('cart', () => {
  const cart = ref<Map<string, OrderItem>>(new Map())

  const items = computed(() : OrderItem[] => {
    let items: OrderItem[] = []
    for (let item of cart.value.values()) {
      console.log(item)
      items.push(item)
    }
    return items
  })

  const count = computed(() => {
    let count = 0;
    for (let item of cart.value.values()) {
      count += item.quantity;
    }
    return count;
  })

  function add(strain: Strain) {
    if (!cart.value.has(strain.id)) {
      cart.value.set(strain.id, new OrderItem(strain, 0))
    }
    let item = cart.value.get(strain.id)
    if (item) {
      item.quantity += 1
    }
  }

  function has(id: string) {
    return cart.value.has(id)
  }

  function set(id: string, quantity: number) {
    if (!cart.value.has(id)) {
      cart.value.set(id, new OrderItem(new Strain(), quantity))
    } else {
      let item = cart.value.get(id)
      if (item) {
        item.quantity = quantity
      }
    }
  }

  function del(id: string) {
    if (cart.value.has(id)) {
      cart.value.delete(id)
    }
  }

  function reset() {
    cart.value.clear()
  }

  return { add, has, set, del, reset, items, count, cart }
})
