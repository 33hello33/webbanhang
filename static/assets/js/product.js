var Vue = new Vue({
  el: '#root',
  delimiters: ['@{', '}'],
  data: {
    product: {id: 0, name: '', unit: '', amount: 0, price: 0, price_import: 0, warehouse: '', id_supplier: 0},
    products: [],
    suppliers: [],
    isUpdate: false,
  },
  mounted(){
    this.listProduct();
    this.listSupplier();
  },
  methods: {
    addProduct() {
      this.isUpdate = false;
      this.changeHeader();
      this.product = {id: 0, name: '', unit: '', amount: 0, price: 0, price_import: 0, warehouse: '', id_supplier: 0};
    },
    listProduct(){
      this.$http.get('product/list').then(response =>{
        if(response.status == 200){
          this.products = response.body;
          this.product = {id: 0, name: '', unit: '', amount: 0, price: 0, price_import: 0, warehouse: '', id_supplier: 0};
        }else{
          console.log(response.body);
        }
      });
    },
    createProduct(product, productIndex){
          this.$http.post('product/create',{
            name: product.name, 
            unit: product.unit, 
            amount: product.amount, 
            price: product.price, 
            price_import: product.price_import, 
            warehouse: product.warehouse, 
            id_supplier: product.id_supplier}).then(response => {
            if(response.status == 200){
                this.listProduct();
            }else{
              console.log(response.body);
            }
          });
    },
    checkForEnter(event){
      if (event.key == "Enter") {
        this.addProduct();
      }
    },
    getDetailProduct(product, productIndex){
      this.isUpdate = true;
      this.changeHeader();
      this.$http.get('product/' + product.id).then(response =>{
        if (response.status == 200){
          this.product = response.body;
        }
      })
    },
    changeHeader(){
      if(this.isUpdate == true){
        this.$refs.headerProduct.innerText = "Sửa Sản Phẩm";
        this.$refs.buttonProduct.innerText = "Sửa Sản Phẩm";
      }else{
        this.$refs.headerProduct.innerText = "Thêm Sản Phẩm";
        this.$refs.buttonProduct.innerText = "Thêm Sản Phẩm";
      }
    },
    deleteProduct(product, productIndex){
      if(confirm("Are you sure ?")){
        this.$http.delete('product/'+ product.id).then(response =>{
          if(response.status == 200){
            this.products.splice(productIndex,1);
            this.product = {id: 0, name: '', unit: '', amount: 0, price: 0, price_import: 0, warehouse: '', id_supplier: 0};
          }
        });
      }
    },
    listSupplier(){
      this.$http.get('supplier/list').then(response =>{
        if(response.status == 200){
          this.suppliers =  response.body;
        }
      });
    },
  }
});