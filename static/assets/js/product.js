const {createApp} = Vue
appProduct = createApp({
  delimiters: ['@{', '}'],
  data() {
    return{
      product: {id: 0, name: '', unit: '', amount: 0, price: 0, price_import: 0, warehouse: '', id_supplier: 0},
      products: [],
      suppliers: [],
      isUpdate: false,
    }
  },
  methods: {
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
        }else{
          console.log(response.data);
        }
      });
    },
    createProduct(product, productIndex){
          axios.post('product/create',{
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
              console.log(response.data);
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