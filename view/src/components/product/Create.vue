<template>
  <div class="container">
    <div class="row">
      <div class="col">
        <form method="post" @submit.prevent="create()">
          <div class="row">
            <div class="mb-3 col-md-6 col-lg-4">
              <label class="form-label" for="product_name">Name</label>
              <input id="product_name" name="Name" class="form-control" v-model="product.Name" maxlength="50" />
            </div>
            <div class="mb-3 col-md-6 col-lg-4">
              <label class="form-label" for="product_price">Price</label>
              <input id="product_price" name="Price" class="form-control" v-model="product.Price" type="number" />
            </div>
            <div class="col-12">
              <router-link class="btn btn-secondary" to="/product">Cancel</router-link>
              <button class="btn btn-primary">Submit</button>
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
  name: 'ProductCreate',
  data() {
    return {
      product: {}
    }
  },
  methods: {
    create() {
      Service.create(this.product).then(() => {
        this.$router.push('/product')
      }).catch((e) => {
        alert(e.response.data)
      })
    }
  }
}
</script>
