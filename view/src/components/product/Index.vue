<template>
  <div class="container">
    <div class="row">
      <div class="col">
        <table class="table table-striped table-hover">
          <thead>
            <tr>
              <th>Id</th>
              <th>Name</th>
              <th>Price</th>
              <th>Actions</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="product in products" :key="product">
              <td class="text-center">{{product.Id}}</td>
              <td>{{product.Name}}</td>
              <td class="text-center">{{product.Price}}</td>
              <td class="text-center">
                <router-link class="btn btn-secondary" :to="`/product/${product.Id}`" title="View"><i class="fa fa-eye"></i></router-link>
                <router-link class="btn btn-primary" :to="`/product/edit/${product.Id}`" title="Edit"><i class="fa fa-pencil"></i></router-link>
                <router-link class="btn btn-danger" :to="`/product/delete/${product.Id}`" title="Delete"><i class="fa fa-times"></i></router-link>
              </td>
            </tr>
          </tbody>
        </table>
        <router-link class="btn btn-primary" to="/product/create">Create</router-link>
      </div>
    </div>
  </div>
</template>
<script>
import Service from './Service'

export default {
  name: 'ProductIndex',
  data() {
    return {
      products: []
    }
  },
  mounted() {
    this.get()
  },
  methods: {
    get() {
      Service.get().then(response => {
        this.products = response.data
      }).catch(e => {
        alert(e.response.data)
      })
    }
  }
}
</script>
