Name:          protoform
Version:       0.5.0
Release:       1%{dist}
Summary:       Protoform - a simple provisioning tool.
License:       MIT
URL:           https://github.com/johndistasio/protoform
Source0:       %{name}-%{version}.tar.gz
BuildRequires: golang

%description
Protoform is a simple provisioning tool intended to set basic properties on a system via templated configuration files to facilitate a more powerful configuration management tool taking over. At it's core, Protoform is nothing more than a template rendering tool.

%define debug_package %{nil}

%prep
%autosetup -n src/github.com/johndistasio/%{name} -c src/github.com/johndistasio/%{name}

%build
make GOPATH=%{_builddir} VERSION=%{version}

%install
install -d %{buildroot}%{_bindir}
install -p -m 0755 ./build/%{name} %{buildroot}%{_bindir}/%{name}

%files
%{_bindir}/%{name}

%changelog
* Sat May 20 2017 John DiStasio <jndistasio@gmail.com> - 0.5.0-1
- Initial version of the package
