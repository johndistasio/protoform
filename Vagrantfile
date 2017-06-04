# -*- mode: ruby -*-
# vi: set ft=ruby :
#
# Virtual machines for build and packaging process development and acceptance
# testing.
#

Vagrant.configure(2) do |config|
  config.vm.define 'cauldron-fedora25' do |guest|
    guest.vm.box = 'bento/fedora-25'
    guest.vm.hostname = 'fedora25'
    guest.vm.provision 'shell', inline: <<-SHELL
      dnf install -y rpm-build make tree golang
    SHELL
  end
end
