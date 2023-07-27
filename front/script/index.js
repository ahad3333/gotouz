const resultContainer = document.getElementById("resultContainer");
const stationFrom = document.querySelector(".stationFrom");
const stationTo = document.querySelector(".stationTo");
// const wagontype = document.querySelector(".wagontype");
const types = document.querySelector(".types");
const time1 = document.querySelector(".time1");
// const time2 = document.querySelector(".time2");
// const time3 = document.querySelector(".time3");
const date = document.querySelector("#date");
const btn = document.querySelector("#btn");
const loadingElement = document.querySelector("#loading");
const cancelBtn = document.querySelector("#cancelBtn");  
const backPath = "http://locahost:8080";

let stopRequests = false;
async function sendRequest(payload) {
  try {
    const response = await fetch(`${backPath}/user`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(payload),
    });

    if (response.ok) {
      const result = await response.json();
      displayResult(result);
      if (!stopRequests) {
        sendRequest(payload);
      }
    } else {
      btn.textContent = "Purchase Ticket"
      btn.style.backgroundColor = "#6751a9";
      console.error("Error:", response.status);
    }
  } catch (error) {
    btn.textContent = "Purchase Ticket"
    btn.style.backgroundColor = "#6751a9";
    console.error("Error:", error);
  }
}

btn.addEventListener("click", (event) => {
  stopRequests = false;
  event.preventDefault();
  btn.textContent = "🔐 stand by..."
  btn.style.backgroundColor = "red";  
  stopBtn.textContent = "Stop";
  const payload = {
    direction: [
      {
        depDate: date.value,
        fullday: true,
        type: "Forward",
      },
    ],
    stationFrom: stationFrom.value,
    stationTo: stationTo.value,
    detailNumPlaces: 1,
    showWithoutPlaces: 0,
    // wagon: {
    //   name: wagontype.value,
    // },
    time: time1.value,
    types: types.value,
  };

  sendRequest(payload);
});

const stopBtn = document.createElement("button");
stopBtn.textContent = "Stop";
stopBtn.style.boxShadow = "0 0 4px 2px red";
stopBtn.style.backgroundColor = "6751a9";
stopBtn.style.width = "200px"
stopBtn.style.borderRadius = "20px"
stopBtn.style.border = "none"
stopBtn.style.fontSize = "20px"
stopBtn.classList.add("button");
stopBtn.addEventListener("click", () => {
  btn.textContent = "Purchase Ticket"
  btn.style.backgroundColor = "#4CAF50";
  stopBtn.textContent = "Stopped 🛑";
  stopRequests = true;
});

// Append the stop button to the document
document.body.appendChild(stopBtn);


function displayResult(result) {
  const resultDiv = document.createElement("div");
  resultDiv.className = "result";
  resultDiv.style.backgroundColor = "#3d5857"
  resultDiv.style.boxShadow = "0 0 2px 2px gray";
  resultDiv.style.marginBottom = "30px";
  resultDiv.style.display = "flex";
  resultDiv.style.flexWrap = "wrap"

  const brand = document.createElement("p");
  brand.textContent = "• Brand: \t" + result.brand;
  brand.style.marginLeft = "20px"
  brand.style.color = "white"
  brand.style.fontSize = "19px";


  const status = document.createElement("p");
  status.style.marginLeft = "15px"
  status.style.color = "white"
  status.style.fontSize = "19px";

  const active = document.createElement("p");
  active.style.marginLeft = "20px"
  active.style.fontSize = "19px";


  const places = document.createElement("p");
  places.textContent = "• Places: \t" + result.places;
  places.style.marginLeft = "20px"
  places.style.color = "white"
  places.style.fontSize = "19px";


  const carnum = document.createElement("p");
  carnum.textContent = "• Carnum: \t" + result.carnum;
  carnum.style.marginLeft = "20px"
  carnum.style.color = "white"
  carnum.style.fontSize = "19px";



  const date = document.createElement("p");
  date.textContent = "• Date: \t" + result.date;
  date.style.marginLeft = "20px"
  date.style.color = "white"
  date.style.fontSize = "19px";


  const now_time = document.createElement("p");
  now_time.textContent = "• NowTime: \t" + result.now_time;
  now_time.style.marginLeft = "20px"
  now_time.style.color = "white"
  now_time.style.fontSize = "19px";


  // const orderId = document.createElement("p");
  // orderId.textContent = "• Order ID: " + result.order_id;
  // orderId.style.marginLeft = "10px"
  // orderId.style.color = "white"


  const getBtn = document.createElement("button");
  getBtn.textContent = "♻️";
  getBtn.style.margin = "10px"
  getBtn.style.width = "70px"
  getBtn.style.borderRadius = "20px"
  getBtn.style.border = "none"
  getBtn.style.backgroundColor = "white";
  getBtn.style.color = "white";
  getBtn.style.fontSize = "20px"
  getBtn.style.right = "20px"
  getBtn.addEventListener("click", async () => {
    try {
      const orderId = result.order_id.replace(/"/g, "");
      const response = await fetch(`https://eticket.railway.uz/api/v1/orders/get/${orderId}`, {
        method: "GET",
        headers: {
          "Content-Type": "application/json",
          "Authorization": "Bearer "+result.token
        },
      });

      if (response.ok) {
        const data = await response.json();
        console.log(data);
        if (`"`+data[0].orderID+`"` == result.order_id) {
          active.style.color = "#74e60a"
          status.textContent = "• Status:";
          active.textContent = "Active "
        }else{
          active.style.color = "red";
          status.textContent = "• Status:";
          active.textContent = "⚠️ NO"
        }
      } else {
        console.error("Error:", response.status);
        active.style.color = "red";
        status.textContent = "• Status:";
        active.textContent = "⚠️ NO"

      }
    } catch (error) {
      active.style.color = "red";
      status.textContent = "• Status:";
      active.textContent = "⚠️ NO"
      console.error("Error:", error);
    }
  });
  const rejectBtn = document.createElement("button");
  rejectBtn.textContent = "🚫"
  rejectBtn.style.backgroundColor = "white";
  rejectBtn.style.margin = "10px"
  rejectBtn.style.width = "70px"
  rejectBtn.style.borderRadius = "20px"
  rejectBtn.style.border = "none"
  rejectBtn.style.fontSize = "20px"
  rejectBtn.addEventListener("click", async () => {
    try {
      const orderId = result.order_id.replace(/"/g, "");
      const cancelPayload = {
        orderId: orderId,
        userId: "user", 
      };
      const response = await fetch(`https://eticket.railway.uz/api/v1/orders/cancel`, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
          "Authorization": "Bearer "+result.token
        },
        body: JSON.stringify(cancelPayload),
      });

      if (response.ok) {
        const data = await response.json();
        active.style.color = "red";
        status.textContent = "• Status:";
        active.textContent = "Cancel";
        console.log(data);
      } else {
        console.error("Error:", response.status);
      }
    } catch (error) {
      active.style.color = "red";
      status.textContent = "• Status:";
      active.textContent = "Cancel";
      console.error("Error:", error);
    }
  });
  const removeBtn = document.createElement("button");
  removeBtn.textContent = "🗑";
  removeBtn.style.backgroundColor = "white";
  // removeBtn.style.color = "white";
  removeBtn.style.margin = "10px"
  removeBtn.style.width = "70px"
  removeBtn.style.borderRadius = "20px"
  removeBtn.style.border = "none"
  removeBtn.style.fontSize = "20px"
  removeBtn.addEventListener("click", () => {
    resultContainer.removeChild(resultDiv);
  });

  resultDiv.appendChild(brand);
  resultDiv.appendChild(carnum);
  resultDiv.appendChild(places);
  resultDiv.appendChild(date);
  resultDiv.appendChild(now_time);
  resultDiv.appendChild(status);
  resultDiv.appendChild(active);
  // resultDiv.appendChild(orderId);
  resultDiv.appendChild(getBtn);
  resultDiv.appendChild(rejectBtn);
  resultDiv.appendChild(removeBtn);

  resultContainer.appendChild(resultDiv);
  setTimeout(() => {
    resultContainer.removeChild(resultDiv);
  }, 12 * 60 * 1000);
}

