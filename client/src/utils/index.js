function priceMod(price) {
  let temp = price
  if (temp % 1 === 0) {
    return `${temp}.000 đ`
  } else {
    return `${temp}00 đ`
  }
}

function productMod(x) {
  const starHTML = `<i class="fas fa-star" style="color:rgb(228, 218, 218)"></i>`
  const yellowStarHTML = `<i class="fas fa-star" style="color:rgb(238, 238, 88)"></i>`
  if (x.desc === '') {
    x.desc = `Không có mô tả.`
  }

  if (!x.rating) {
    x.rating = `Hiện chưa có đánh giá.<br><br>`
  } else {
    let temp = ''
    for (let i = 0; i < Math.round(x.rating); i++) {
      temp += yellowStarHTML
    }
    for (let i = 0; i < 5 - Math.round(x.rating); i++) {
      temp += starHTML
    }
    x.rating = temp + `<br><div>(${x.numOfRate} lượt đánh giá)</div>`
  }
  x.price = priceMod(x.price)
}

export { priceMod, productMod }
