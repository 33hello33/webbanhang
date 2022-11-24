appProduct = createApp({
  delimiters: ['@{', '}'],
  data() {
    return{
      product: {id: 0, name: '', unit: '', amount: 0, price: 0, price_import: 0, warehouse: '', barcode: ''},
      products: [],
      isUpdate: false,
      timeOut: 0,
      searchInput: '',
      file:'',
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
      this.product = {id: 0, name: '', unit: '', amount: 0, price: 0, price_import: 0, warehouse: '', barcode: ''};
    },
    listProduct(){
      axios.get('product/list').then(response =>{
        if(response.status == 200){
          this.products = response.data;
          this.product = {id: 0, name: '', unit: '', amount: 0, price: 0, price_import: 0, warehouse: '', barcode: ''};
        }
      })
      .catch(error => {
        console.log(error.data.Error);
        });   
    },
    createProduct(product){
      if(this.isUpdate == false){ // create product
        axios.post('product/create',product)
        .then(response => {
          if(response.status == 200){
              this.listProduct();
          }
        })
        .catch(error => {
          console.log(error.data.Error);
          });  
      }else{ // update product
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
    getDetailProduct(product){
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
        axios.delete('product/'+ product.id)
        .then(response =>{
          if(response.status == 200){
            this.products.splice(productIndex,1);
            this.product = {id: 0, name: '', unit: '', amount: 0, price: 0, price_import: 0, warehouse: '', barcode: ''};
          }
        })
        .catch(error =>{
          if(error.data.Error.includes('foreign key')){
            alert("Không thể xóa sản phẩm vì sản phẩm này đang tồn tại trong hóa đơn");
          }else{
            console.log(error.data.Error);
          }
        });
      }
    },
    copyProduct(product){
       // create product
       axios.post('product/copy/' +product.id , product)
       .then(response => {
         if(response.status == 200){
             this.listProduct();
         }
       })
       .catch(error => {
         console.log(error.data.Error);
         });  
    },
    uploadFile(){
      let formData = new FormData();
      console.log(this.file);
      formData.append('file', this.file);
      axios.post('product/import_from_file',
        formData,
        {
          headers:{'Content-Type':'multipart/form-data'}
        })
        .then(response =>{
          console.log(response.data);
        })
        .catch(error =>{
          console.log(error.data);
        })
    },
    inputFileChange(){
      this.file=this.$refs.file.files[0];
    },
    exportFile(){
      const url = 'http://localhost:8080/product/export_to_file';
        axios.get('product/export_to_file',
          {responseType : 'blob'})
        .then(({data}) => {
          //var obj = JSON.parse(data);
          //console.log(obj);
          const downloadUrl = window.URL.createObjectURL(new Blob([data]));
          const link = document.createElement('a');
          link.href = downloadUrl;
          link.setAttribute('download', 'file.csv'); //any other extension
          document.body.appendChild(link);
          link.click();
          link.remove();
        }); 
      },
      autoGenBarcode(){
        var m = new Date().valueOf();
        let i=0;
        for(;i < 100; i++){
          if (this.genBarcode(m+i) == true){
              break
          }
        }
        this.product.barcode = (m +i).toString();
      },
      genBarcode(number){
        try {
          JsBarcode("#barcode")
          .options({width:1.5,height:50,})
          .EAN13(number, {fontSize: 12, textMargin: 0})
          .render();
          return true
        }
        catch(err) {
          return false
        }
      }
  },
  beforeMount(){
    this.listProduct();
  },
});
appProduct.mount("#root")