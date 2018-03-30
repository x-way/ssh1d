# Maintainer: Andreas Jaggi <andreas.jaggi@waterwave.ch>

pkgname=ssh1d
pkgver=0.1
pkgrel=1
pkgdesc="ssh1d - fake SSH-1 daemon"
arch=('x86_64' 'i686')
url="http://github.com/x-way/$pkgname/"
license=('MIT')
depends=('systemd')
makedepends=('go')
options=('!strip' '!emptydirs')
source=("http://github.com/x-way/$pkgname/$pkgname-$pkgver.tar.gz")
sha256sums=('bb37678fbbf3ea458e2e3d21c6cacdd1e8e04c3e9767091988e390b32b3fa2bb')

prepare() {
  cd "$pkgname-$pkgver"

  GOPATH="`pwd`" go get -d -v
}

build() {
  cd "$pkgname-$pkgver"

  GOPATH="`pwd`" go build
}

package() {
  cd "$pkgname-$pkgver"

  install -Dm755 "$pkgname-$pkgver" "$pkgdir/usr/bin/$pkgname"
  install -Dm644 LICENSE "$pkgdir/usr/share/licenses/$pkgname/LICENSE"
	install -Dm644 "$pkgname.socket" "$pkgdir/usr/lib/systemd/system/$pkgname.socket"
	install -Dm644 "$pkgname.service" "$pkgdir/usr/lib/systemd/system/$pkgname.service"
}

# vim:set ts=2 sw=2 et:
