package database

import "errors"

var (
	ErrCantFindProduct    = errors.New("can't find product")
	ErrCantDecodeProducts = errors.New("can't decode product")
	ErrUserIdIsNotValid   = errors.New("This user is invalid")
	ErrCantUpdateUser     = errors.New("cannot add this product to the cart")
	ErrCantRemoveItemCart = errors.New("cannot remove this item from cart")
	ErrCantGetItem        = errors.New("was unable to get the item from the cart")
	ErrCantBuyCartItem    = errors.New("can not update the purchase")
)

func AddProductToCart() {}
func RemoveCartItem()   {}
func BuyItemFromCart()  {}
func InstantBuyer()     {}
