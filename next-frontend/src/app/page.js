"use client"

import { useRouter } from 'next/navigation';

export default function Login() {
  const router = useRouter();

  useEffect(() => {
    router.push('/login');
  });
  
  return null;
}
