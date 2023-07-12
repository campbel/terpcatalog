import { ref, computed, watch } from 'vue'
import { defineStore } from 'pinia'
import { OrderItem } from '@/types/order'
import { Strain } from '@/types/strain'

export const useCartStore = defineStore('cart', () => {
  const cart = ref<Map<string, OrderItem>>(new Map())

  // Load cart from local storage
  let loadedCart = localStorage.getItem('cart')
  if (loadedCart) {
    JSON.parse(loadedCart)?.forEach((item: any) => {
      cart.value.set(item[0], new OrderItem(item[1].strain, item[1].quantity))
    })
  }

  // Save cart to local storage
  watch(cart, (newCart) => {
    localStorage.setItem('cart', JSON.stringify(Array.from(newCart.entries())))
  }, { deep: true })

  const items = computed((): OrderItem[] => {
    let items: OrderItem[] = []
    for (let item of cart.value.values()) {
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

  const total = computed(() => {
    let total = 0;
    for (let item of cart.value.values()) {
      total += item.quantity * item.strain.price;
    }
    return total;
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

  function get(id: string): OrderItem {
    if (cart.value.has(id)) {
      let item = cart.value.get(id)
      if (item) {
        return item
      }
    }
    return new OrderItem(new Strain(), 0)
  }

  function del(id: string) {
    if (cart.value.has(id)) {
      cart.value.delete(id)
    }
  }

  function reset() {
    cart.value.clear()
  }

  return { add, has, set, get, del, reset, items, count, total }
})
