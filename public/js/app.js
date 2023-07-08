
const { createApp } = Vue

createApp({
    data() {
        return {
            loading: true,
            catalog: []
        }
    },
    created() {
        this.getCatalog()
    },

    methods: {
        async getCatalog() {
            try {
                this.loading = true
                const res = await fetch('/api/catalog')
                const data = await res.json()
                this.catalog = data
                console.log(this.catalog)
                this.loading = false
            } catch (error) {
                console.log(error)
            }
        }
    },
}).mount('#app')