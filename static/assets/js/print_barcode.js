const {createApp} = Vue 
appPrintBarcode = createApp({
    delimiters: ['@{', '}'],
    data(){
      return {
      }
    },
    methods: {
      print(){
        const url = window.location.href;
        const amount = url.split("/").slice(-1)[0];
        const barcode = url.split("/").slice(-2)[0];

        const barcode_value = barcode.split("=")[1] ;
        const amount_value = amount.split("=")[1] ;

        // have 1 div original, so just need to copy (amount-1)
        for(let i=1;i < amount_value;i++){
          var showBarcode = document.getElementsByClassName("div_barcode");
          var clone = showBarcode[0].cloneNode(true);
          showBarcode[0].parentNode.appendChild(clone);
        }
       
        JsBarcode(".barcode")
        .options({width:1.5,height:50,})
        .EAN13(barcode_value, {fontSize: 12, textMargin: 0})
        .render();
      }
    },
    mounted(){
      this.print();
    }
  });
appPrintBarcode.mount("#root")