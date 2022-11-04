appProduct = createApp({
  delimiters: ['@{', '}'],
  data() {
    return{
      product: {id: 0, name: '', unit: '', amount: 0, price: 0, price_import: 0, warehouse: '', id_supplier: 0},
      products: [],
      suppliers: [],
      isUpdate: false,
      timeOut: 0,
      searchInput: '',
    }
  },
  methods: {
    searchProduct(){
      if(this.searchInput == ""){
        this.listProduct();
        return;
      }
      axios.get('product/search/'+this.searchInput)
      .then(response=>{
        if(response.status == 200){
          this.products = response.data;
        }
      })
      .catch(error => {
        console.log('get list product err: ' + error.data.Error);
      });   
    },
    searchInputChanges(){
      clearTimeout(this.timeOut);

      this.timeOut = setTimeout(() => {
        this.searchProduct();
      }, 300);
    },
    addProduct() {
      this.isUpdate = false;
      this.changeHeader();
      this.product = {id: 0, name: '', unit: '', amount: 0, price: 0, price_import: 0, warehouse: '', id_supplier: 0};
    },
    listProduct(){
      axios.get('product/list').then(response =>{
        if(response.status == 200){
          this.products = response.data;
          this.product = {id: 0, name: '', unit: '', amount: 0, price: 0, price_import: 0, warehouse: '', id_supplier: 0};
        }
      })
      .catch(error => {
        alert(error.data.Error);
        });   
    },
    createProduct(product, productIndex){
      if(this.isUpdate == false){ // create product
        axios.post('product/create',product)
        .then(response => {
          if(response.status == 200){
              this.listProduct();
          }
        })
        .catch(error => {
          alert(error.data.Error);
          });   
      }
      else{ // update product
        axios.put('product/update', product)
        .then(  response => {
          if (response.status == 200){
          this.listProduct();
        }
      })
      }
    },
    checkForEnter(event){
      if (event.key == "Enter") {
        this.addProduct();
      }
    },
    getDetailProduct(product, productIndex){
      this.isUpdate = true;
      this.changeHeader();
      axios.get('product/' + product.id).then(response =>{
        if (response.status == 200){
          this.product = response.data;
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
        axios.delete('product/'+ product.id).then(response =>{
          if(response.status == 200){
            this.products.splice(productIndex,1);
            this.product = {id: 0, name: '', unit: '', amount: 0, price: 0, price_import: 0, warehouse: '', id_supplier: 0};
          }
        });
      }
    },
    listSupplier(){
      axios.get('supplier/list').then(response =>{
        if(response.status == 200){
          this.suppliers =  response.data;
        }
      });
    },
  },
  beforeMount(){
    this.listProduct();
    this.listSupplier();
  },
});
appProduct.mount("#root")