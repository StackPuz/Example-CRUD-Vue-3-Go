<template>
  <div class="container">
    <div class="row">
      <div class="col">
        <form method="post" @submit.prevent="this.delete()">
          <div class="row">
            <div class="mb-3 col-md-6 col-lg-4">
              <label class="form-label" for="product_id">Id</label>
              <input readonly id="product_id" name="Id" class="form-control" :value="product.Id" type="number" required />
            </div>
            <div class="mb-3 col-md-6 col-lg-4">
              <label class="form-label" for="product_name">Name</label>
              <input readonly id="product_name" name="Name" class="form-control" :value="product.Name" maxlength="50" />
            </div>
            <div class="mb-3 col-md-6 col-lg-4">
              <label class="form-label" for="product_price">Price</label>
              <input readonly id="product_price" name="Price" class="form-control" :value="product.Price" type="number" />
            </div>
            <div class="col-12">
              <router-link class="btn btn-secondary" to="/product">Cancel</router-link>
              <button class="btn btn-danger">Delete</button>
            </div>
          </div>
        </form>
      </div>
    </div>
  </div>
</template>
<script>
import Service from './Service'

export default {
  name: 'ProductDelete',
  data() {
    return {
      product: {}
    }
  },
  mounted() {
    this.get()
  },
  methods: {
    get() {
      return Service.delete(this.$route.params.id).then(response => {
        this.product = response.data
      })
    },
    delete() {
      Service.delete(this.$route.params.id, this.product).then(() => {
        this.$router.push('/product')
      }).catch((e) => {
        alert(e.response.data)
      })
    }
  }
}
</script>
