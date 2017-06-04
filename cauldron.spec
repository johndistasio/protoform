Name:          cauldron
Version:       0.7.1
Release:       1%{dist}
Summary:       Cauldron - a simple provisioning tool.
License:       MIT
URL:           https://github.com/johndistasio/cauldron
Source0:       %{name}-%{version}.tar.gz
BuildRequires: golang

%description
Cauldron is a simple provisioning tool intended to set basic properties on a system via templated configuration files to facilitate a more powerful configuration management tool taking over.

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
* Sat Jun 3 2017 John DiStasio <jndistasio@gmail.com> - 0.7.1-1
- Increment version to 0.7.1
* Sat Jun 3 2017 John DiStasio <jndistasio@gmail.com> - 0.6.0-1
- Increment version to 0.6.0
* Sat May 20 2017 John DiStasio <jndistasio@gmail.com> - 0.5.0-1
- Initial version of the package
