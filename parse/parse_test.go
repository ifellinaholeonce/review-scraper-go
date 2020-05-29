package Parse

import (
	"fmt"
	"strings"
	"testing"
)

var shopifyReviewHtml = `
<div class="review-listing ">
  <div data-review-id="123456">
  <div class="review-listing-header">
    <h3 class="review-listing-header__text">
      My Store
    </h3>
  </div>
  <div class="review-metadata">
    <div class="review-metadata__item">
      <div class="review-metadata__item-label">
        Rating
      </div>
      <div class="review-metadata__item-value">
        <div class="ui-star-rating" data-rating="4"><div class="ui-star-rating__icon-set" aria-hidden="true"><div class="ui-star-rating__icon"><svg class="icon" aria-hidden="true" focusable="false"> <use xlink:href="#v2-icons-ui-star"></use> </svg></div><div class="ui-star-rating__icon"><svg class="icon" aria-hidden="true" focusable="false"> <use xlink:href="#v2-icons-ui-star"></use> </svg></div><div class="ui-star-rating__icon"><svg class="icon" aria-hidden="true" focusable="false"> <use xlink:href="#v2-icons-ui-star"></use> </svg></div><div class="ui-star-rating__icon"><svg class="icon" aria-hidden="true" focusable="false"> <use xlink:href="#v2-icons-ui-star"></use> </svg></div><div class="ui-star-rating__icon"><svg class="icon" aria-hidden="true" focusable="false"> <use xlink:href="#v2-icons-ui-star"></use> </svg></div></div><div class="ui-star-rating__text"><span class="ui-star-rating__rating visuallyhidden">4 of 5 stars</span></div></div>
      </div>
    </div>
    <div class="review-metadata__item">
      <div class="review-metadata__item-label">
        Posted
      </div>
      <div class="review-metadata__item-value">
        May 22, 2020
      </div>
    </div>
  </div>
  <div class="review-content" data-truncate-review="">
    <div class="truncate-content-copy">
      <p>I've loved using Smile so far. It's easy to install and use. I love the amazing support they offer like tips on their blog. I've also received speedy help from the Team since I installed. Maggie was a gem to work with.</p>
    </div>
    <button name="button" type="submit" class="marketing-button--visual-link truncate-content-toggle">Show full review</button>
  </div>
  </div>
  <div class="review-footer">

    <div data-selected="false" data-disabled="false" class="review-helpfulness">
      <form action="/reviews/558902/helpful" accept-charset="UTF-8" method="post"><input name="utf8" type="hidden" value="✓"><input type="hidden" name="authenticity_token" value="NeIe20z5Vkr/f3w5OQXtaPejaZth4HaljCBqPfqSRUbMfunkUlV8v43ShpV88Ib8Y/eNuElLmhpbB0YCSuXeng==">
        <button name="button" type="submit" class="review-action__button review-action__button-helpfulness" aria-label="Log in to vote this review as helpful.">
          Helpful
          (<span class="review-helpfulness__helpful-count">0</span>)
</button>
        <span class="review-helpfulness__helpful-flash-message" aria-live="polite"></span>
</form></div>  </div>
</div>
`

func Test_parse(t *testing.T) {
	shopifyReviewHtml = shopifyReviewHtml + shopifyReviewHtml
	result, error := Parse(strings.NewReader(shopifyReviewHtml))
	if error != nil {
		t.Errorf("Parse error")
	}
	if result == nil {
		t.Errorf("Parse was incorrect, got: %v", nil)
	}
	if len(result) != 2 {
		t.Errorf("Got length %v, expected: 2", len(result))
	}
	fmt.Println(result)
}
