//Ajax solo funciona en servidores por el protocolo http
text="";inputAux="";
textClick();
function textClick(){
    if(inputAux){
        inputAux.value = text;
        //autocomplete.innerHTML='';
        return inputAux.form;
    }
}
function select(value){
    text=value;
    textClick().submit();
}
function selected(value){
    text=value;
    textClick();
    document.querySelector(".autocom-box").remove();
}
function autocom(inputBox,rute,click){
    if(inputBox){
        //getting all required elements
        let currentFocus=0,varx=[];
        //create autocomplete
        function insertAfter(e,i){ 
            if(e.nextSibling){ 
                e.parentNode.insertBefore(i,e.nextSibling); 
            } else { 
                e.parentNode.appendChild(i); 
            }
        }
        divcom=document.createElement("div");
        divcom.setAttribute('class','autocom-box');
        insertAfter(inputBox,divcom);
        autocomplete = document.querySelector(".autocom-box");
        autocomplete.remove();
        //Events[autocomple]
        inputBox.addEventListener("keyup",function (e){
            //getting all required elements
            var dataAJax,datas,listData;
            //Peticcion Ajax
            const XHR = new XMLHttpRequest();
            //Obtencion de respuesta de la otra pagina
            XHR.open('POST',rute+"/"+inputBox.value,true);
            XHR.send(0);
            
            XHR.onreadystatechange = function() {
                if(this.readyState == "4" && this.status =="200" ){
                    //Requesting data
                    var charCode = e.keyCode;
                    if(charCode!=13&&charCode!=40&&charCode!=38){
                        autocomplete.innerHTML='';
                        dataAJax = JSON.parse(this.responseText);
                        for(let i = 0; i < 6; i++){
                            datas=dataAJax[i];//optional chaining operator (?.)
                            if(inputBox.value.length>0&&datas!==undefined){
                                listData = `<div onclick="${(click=="click")?'select(this.innerHTML)':'selected(this.innerHTML)'}">${datas}</div>`;
                                autocomplete.insertAdjacentHTML("afterbegin", listData);
                            }
                        }
                        if(inputBox.value.length==0||!autocomplete.querySelector("div")){
                            //autocomplete.innerHTML='';
                            autocomplete.remove();
                        }else{
                            insertAfter(inputBox,divcom);
                        }
                        if(currentFocus==0&&charCode!=38)varx=[inputBox.value];
                    }
                }
            }
        });
        inputBox.addEventListener("keydown",function(e){
            var charCode = e.keyCode;
            let eraser=autocomplete.querySelectorAll('div').length;
            if(eraser!=0){
                //Eventos de deslizamiento
                if(charCode == 40){
                    /*If the arrow DOWN key is pressed,
                    increase the currentFocus inputBox.value:*/
                    currentFocus++;
                    /*and and make the current item more visible:*/
                }else if(charCode == 38){ //up
                    /*If the arrow UP key is pressed,
                    decrease the currentFocus inputBox.value:*/
                    currentFocus--;
                    /*and and make the current item more visible:*/
                }else{
                    currentFocus=0;
                }
                if(currentFocus < 0){
                    currentFocus=eraser;
                }
                if(currentFocus > eraser+1){
                    currentFocus=1;
                }
                if(charCode==13&&click==""){e.preventDefault();autocomplete.remove();}
                let x = autocomplete.querySelector(".autocom-box div:nth-child("+currentFocus+")");
                x?.setAttribute("class","autocomplete-active");
                if(currentFocus>1)autocomplete.querySelector(".autocom-box div:nth-child("+(currentFocus-1)+")").removeAttribute("class");
                if(currentFocus<eraser)autocomplete.querySelector(".autocom-box div:nth-child("+(currentFocus+1)+")").removeAttribute("class");
                if(charCode==40||charCode==38){
                inputBox.value=(currentFocus>0&&currentFocus<eraser+1)?x.innerHTML:varx;
                e.preventDefault();
                moveCursorToEnd();
                }
            }
        });
        document.addEventListener("mouseup", function(e) {
            if (!inputBox.contains(e.target)&&!autocomplete.contains(e.target)) {
                autocomplete.remove();
            }
        });
        function moveCursorToEnd() {
            setTimeout(function(){
                inputBox.selectionStart = inputBox.selectionEnd = inputBox.value.length;
                inputBox.focus();
            },0);
        }
        inputAux=inputBox;
    }
}
