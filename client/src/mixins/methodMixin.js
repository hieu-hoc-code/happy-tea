var methodMixins = {
  methods: {
    goDetailPage(e) {
      const id = e.target.closest('.card').dataset.id
      this.$router.push({ name: 'detail', params: { id: id } })
    },
    descriptionMod(x) {
      if (x === '') {
        return `Không có mô tả.`
      }
      return x
    },
    ratingMod(rating, numOfRate) {
      const starHTML = `<i class="fas fa-star" style="color:rgb(228, 218, 218)"></i>`
      const yellowStarHTML = `<i class="fas fa-star" style="color:rgb(238, 238, 88)"></i>`
  
      if (!rating) {
        rating = `Hiện chưa có đánh giá.<br><br>`
      } else {
        let temp = ''
        for (let i = 0; i < Math.round(rating); i++) {
          temp += yellowStarHTML
        }
        for (let i = 0; i < 5 - Math.round(rating); i++) {
          temp += starHTML
        }
        rating = temp + `<br><div>(${numOfRate} lượt đánh giá)</div>`
      }
      return rating
    },
    priceMod(price) {
      let temp = price
      if (temp % 1 === 0) {
        return `${temp}.000 đ`
      } else {
        return `${temp}00 đ`
      }
    },
  },
}
export default methodMixins
