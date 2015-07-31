# This is just an example for a Brew formula
# If you want to create a private tap just place this in a repo name homebrew-<name>

class Tf < Formula
  homepage "https://github.com/mickep76/go-sampleapp"
  url "https://github.com/mickep76/go-sampleapp/archve/0.1.tar.gz"
  sha256 "c9e6f021f135555e89b2da25f2c5b3a8ec77e3eaa48ebbb8a95091729e485711"

  depends_on "go" => :build

  def install
    ENV["GOPATH"] = buildpath
    system "./build"
    bin.install "bin/go-sampleapp"
  end

  test do
    system "#{bin}/go-sampleapp", "--version"
  end
end
