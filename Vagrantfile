# -*- mode: ruby -*-
# vi: set ft=ruby :
#
# Virtual machines for build and packaging process development and acceptance
# testing.
#

RPM = "yum -y install rpm-build make golang"

Vagrant.configure(2) do |config|
  config.vm.define 'protoform-fedora25' do |guest|
    guest.vm.box = 'bento/fedora-25'
    guest.vm.hostname = 'fedora25'
    guest.vm.provision 'shell', inline: RPM
  end

  config.vm.define 'protoform-centos73' do |guest|
    guest.vm.box = 'bento/centos-7.3'
    guest.vm.hostname = 'centos73'
    guest.vm.provision 'shell', inline: RPM
  end
end
